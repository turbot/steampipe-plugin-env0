package env0

import (
	"context"

	"github.com/env0/terraform-provider-env0/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "organization_id",
			Description: "Unique identifier for the organization.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getOrganizationId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getOrganizationMemoized = plugin.HydrateFunc(getOrganizationUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getOrganizationId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org, err := getOrganizationMemoized(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationId", "error", err)
		return nil, err
	}
	return org.(client.Organization).Id, nil
}

// Build a cache key for the call to getOrganizationCacheKey.
func getOrganizationCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getOrganization"
	return key, nil
}

func getOrganizationUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationUncached", "connection_error", err)
		return nil, err
	}

	organization, err := client.Organization()
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationUncached", "api_error", err)
		return nil, err
	}
	return organization, nil
}
