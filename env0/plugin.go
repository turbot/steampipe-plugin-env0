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
			"env0_organization": tablelaunchdarklyAccessToken(ctx),
			// "env0_account_member": tableenv0AccountMember(ctx),
			// "env0_audit_log":      tableenv0AuditLog(ctx),
			// "env0_environment":    tableenv0Environment(ctx),
			// "env0_feature_flag":   tableenv0FeatureFlag(ctx),
			"env0_project": tableenv0Project(ctx),
			// "env0_team":           tableenv0Team(ctx),
		},
	}
	return p
}
