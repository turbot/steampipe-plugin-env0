package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0APIKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_api_key",
		Description: "Returns information about the env0 API keys.",
		List: &plugin.ListConfig{
			Hydrate: listAPIKeys,
		},

		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the API key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier for the API key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_key_id",
				Description: "API Key ID value.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "The organization ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_role",
				Description: "Associated organization role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_used_at",
				Description: "Returns the time and date when the API key was last used.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastUsedAt").NullIfZero(),
			},
			{
				Name:        "created_at",
				Description: "Returns the time and date when the API key was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by",
				Description: "Name of the authentication type used for creation of API key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by_user",
				Description: "Details of the user who created the API key.",
				Type:        proto.ColumnType_JSON,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listAPIKeys(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_api_key.listAPIKeys", "connection_error", err)
		return nil, err
	}

	apiKeys, err := client.ApiKeys()
	plugin.Logger(ctx).Error("API KEY LENGTH:", len(apiKeys))
	if err != nil {
		logger.Error("env0_api_key.listAPIKeys", "api_error", err)
		return nil, err
	}

	for _, apiKey := range apiKeys {
		d.StreamListItem(ctx, apiKey)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
