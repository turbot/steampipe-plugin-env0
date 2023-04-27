package env0

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type env0Config struct {
	APIKey    *string `cty:"api_key"`
	APISecret *string `cty:"api_secret"`
	OrganizationId *string `cty:"organization_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"api_secret": {
		Type: schema.TypeString,
	},
	"organization_id": {
		Type: schema.TypeString,
	},
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
