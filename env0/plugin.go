package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (env0) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-env0",
		DefaultTransform: transform.FromCamel(),
		// DefaultIgnoreConfig: &plugin.IgnoreConfig{
		// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		// },
		// DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"env0_api_key":      tableEnv0APIKey(ctx),
			"env0_environment":  tableenv0Environment(ctx),
			"env0_notification": tableEnv0Notification(ctx),
			"env0_organization": tableEnv0Organization(ctx),
			"env0_project":      tableenv0Project(ctx),
			"env0_role":         tableEnv0Role(ctx),
			"env0_team":         tableEnv0Team(ctx),
			"env0_template":     tableenv0Template(ctx),
			"env0_user":         tableEnv0User(ctx),
		},
	}
	return p
}
