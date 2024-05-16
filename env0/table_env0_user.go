package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0User(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_user",
		Description: "Returns information about the env0 users.",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
			{
				Name:        "email",
				Description: "The email ID of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Email"),
			},
			{
				Name:        "family_name",
				Description: "The family name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.FamilyName"),
			},
			{
				Name:        "given_name",
				Description: "The given name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.GivenName"),
			},
			{
				Name:        "user_id",
				Description: "A unique identifier of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.UserId"),
			},
			{
				Name:        "picture",
				Description: "The picture of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Picture"),
			},
			{
				Name:        "created_at",
				Description: "Time when the user was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("User.CreatedAt"),
			},
			{
				Name:        "last_login",
				Description: "Time when the user last logged in.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("User.LastLogin"),
			},
			{
				Name:        "app_metadata",
				Description: "User's app metadata.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User.AppMetadata"),
			},
			{
				Name:        "role",
				Description: "Assigned role of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "Status of the user.",
				Type:        proto.ColumnType_STRING,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_user.listUsers", "connection_error", err)
		return nil, err
	}

	users, err := client.Users()
	if err != nil {
		logger.Error("env0_organenv0_userization.listUsers", "api_error", err)
		return nil, err
	}

	for _, user := range users {
		d.StreamListItem(ctx, user)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
