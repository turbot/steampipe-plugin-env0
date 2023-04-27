package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0Team(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_team",
		Description: "Returns information about the Env0 teams.",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "Team ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "Organization ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A brief description of the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "users",
				Description: "List the users in a team.",
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

func listTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_team.listTeams", "connection_error", err)
		return nil, err
	}

	teams, err := client.Teams()
	if err != nil {
		logger.Error("env0_organization.listOrganizations", "api_error", err)
		return nil, err
	}

	for _, team := range teams {
		d.StreamListItem(ctx, team)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_team.getTeam", "connection_error", err)
		return nil, err
	}

	team, err := client.Team(id)
	if err != nil {
		logger.Error("env0_team.getTeam", "api_error", err)
		return nil, err
	}

	return team, nil
}
