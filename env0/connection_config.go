package env0

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type env0Config struct {
	APIKey    *string `hcl:"api_key"`
	APISecret *string `hcl:"api_secret"`
}

func ConfigInstance() interface{} {
	return &env0Config{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) env0Config {
	if connection == nil || connection.Config == nil {
		return env0Config{}
	}
	config, _ := connection.Config.(env0Config)
	return config
}
