package env0

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	env0CLient "github.com/env0/terraform-provider-env0/client"
	"github.com/env0/terraform-provider-env0/client/http"
	"github.com/go-resty/resty/v2"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (env0CLient.ApiClientInterface, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "env0"

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(env0CLient.ApiClientInterface), nil
	}

	api_key := os.Getenv("ENV0_API_KEY")
	api_secret := os.Getenv("ENV0_API_SECRET")

	env0Config := GetConfig(d.Connection)

	if env0Config.APIKey != nil {
		api_key = *env0Config.APIKey
	}
	if env0Config.APISecret != nil {
		api_secret = *env0Config.APISecret
	}

	// Check for env0 CLI configuration

	if api_key == "" && api_secret == "" {
		home, _ := os.UserHomeDir()
		file, _ := os.ReadFile(home + "/.env0/config.json")
		cliCreds := make(map[string]string)

		_ = json.Unmarshal([]byte(file), &cliCreds)

		for k, v := range cliCreds {
			if k == "apiKey" {
				api_key = v
			}
			if k == "apiSecret" {
				api_secret = v
			}
		}
	}

	if api_key == "" {
		// api_key not set
		return nil, errors.New("api_key must be configured")
	}
	if api_secret == "" {
		// api_secret not set
		return nil, errors.New("api_secret must be configured")
	}

	httpClient, err := http.NewHttpClient(http.HttpClientConfig{
		ApiKey:      api_key,
		ApiSecret:   api_secret,
		ApiEndpoint: "https://api.env0.com/",
		UserAgent:   "",
		RestClient:  resty.New(),
	})
	if err != nil {
		plugin.Logger(ctx).Error("Error in creating the client", err)
	}
	client := env0CLient.NewApiClient(httpClient, "")

	d.ConnectionManager.Cache.Set(cacheKey, client)
	return client, nil
}
