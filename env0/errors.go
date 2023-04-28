package env0

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for LaunchDarkly API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range notFoundErrors {
			// handle not found error
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}