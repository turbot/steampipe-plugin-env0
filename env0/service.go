package env0

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

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

	// Default to using env vars (#2)
	api_key := os.Getenv("ENV0_API_KEY")
	api_secret := os.Getenv("ENV0_API_SECRET")
	organization_id := os.Getenv("ENV0_ORGANIZATION_ID")

	// But prefer the config (#1)
	env0Config := GetConfig(d.Connection)

	if env0Config.APIKey != nil {
		api_key = *env0Config.APIKey
	}
	if env0Config.APISecret != nil {
		api_secret = *env0Config.APISecret
	}
	if env0Config.OrganizationId != nil {
		organization_id = *env0Config.OrganizationId
	}

	if api_key == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}
	if api_secret == "" {
		// Credentials not set
		return nil, errors.New("api_secret must be configured")
	}
	if organization_id == "" {
		// Credentials not set
		return nil, errors.New("organization_id must be configured")
	}

	apiEndpoint := "https://api.env0.com/"
	// organizationId := "fe8c0666-b2ed-498f-a94b-3f57e1923021"

	isIntegrationTest := false

	restyClient := resty.New()

	restyClient.
		SetRetryCount(5).
		SetRetryWaitTime(time.Second).
		SetRetryMaxWaitTime(time.Second * 5).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			if r == nil {
				// No response. Possiblly a networking issue (E.g. DNS lookup failure).
				log.Printf("[WARN] No response, retrying request: %s %s", r.Request.Method, r.Request.URL)
				return true
			}

			// When running integration tests 404 may occur due to "database eventual consistency".
			// Retry when there's a 5xx error. Otherwise do not retry.
			if r.StatusCode() >= 500 || isIntegrationTest && r.StatusCode() == 404 {
				log.Printf("[WARN] Received %d status code, retrying request: %s %s", r.StatusCode(), r.Request.Method, r.Request.URL)
				return true
			}

			if r.StatusCode() == 200 && isIntegrationTest && r.String() == "[]" {
				log.Printf("[WARN] Received an empty list for an integration test, retrying request: %s %s", r.Request.Method, r.Request.URL)
				return true
			}

			return false
		})

	httpClient, err := http.NewHttpClient(http.HttpClientConfig{
		ApiKey:      api_key,
		ApiSecret:   api_secret,
		ApiEndpoint: apiEndpoint,
		UserAgent:   "",
		RestClient:  restyClient,
	})
	if err != nil {
		plugin.Logger(ctx).Error("Error in creating the client", err)
	}
	client := env0CLient.NewApiClient(httpClient, organization_id)

	d.ConnectionManager.Cache.Set(cacheKey, client)
	return client, nil
}