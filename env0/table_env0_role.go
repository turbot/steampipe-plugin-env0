package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0Role(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_role",
		Description: "Returns information about the env0 roles.",
		List: &plugin.ListConfig{
			Hydrate: listRoles,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRole,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "The organization ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "permissions",
				Description: "Permissions associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "is_default_role",
				Description: "Returns tru if the role is default.",
				Type:        proto.ColumnType_BOOL,
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

func listRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_role.listRoles", "connection_error", err)
		return nil, err
	}

	roles, err := client.Roles()
	if err != nil {
		logger.Error("env0_role.listRoles", "api_error", err)
		return nil, err
	}

	for _, role := range roles {
		d.StreamListItem(ctx, role)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_role.getRole", "connection_error", err)
		return nil, err
	}

	role, err := client.Role(id)
	if err != nil {
		logger.Error("env0_role.getRole", "api_error", err)
		return nil, err
	}

	return role, nil
}
