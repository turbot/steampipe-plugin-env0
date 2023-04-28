package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0Notification(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_notification",
		Description: "Returns information about the env0 notifications.",
		List: &plugin.ListConfig{
			Hydrate: listNotification,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the notification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "Notification ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "Organization ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "value",
				Description: "Value of notification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "Name of the authentication type used for creation of authentication.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by_user",
				Description: "Details of the user who created the notification.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "type",
				Description: "Type of notification.",
				Type:        proto.ColumnType_STRING,
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

func listNotification(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_notification.listNotification", "connection_error", err)
		return nil, err
	}

	notifications, err := client.Notifications()
	if err != nil {
		logger.Error("env0_notification.listNotification", "api_error", err)
		return nil, err
	}

	for _, notification := range notifications {
		d.StreamListItem(ctx, notification)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
