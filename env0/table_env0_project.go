package env0

import (
	"context"

	"github.com/env0/terraform-provider-env0/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableenv0Project(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_project",
		Description: "Returns information about the env0 project.",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "Unique identifier of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "The organization ID in which the project is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_project_id",
				Description: "Parent project id of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "Name of the authentication type used for creation of project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by_user",
				Description: "Details of the user who created the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description associated with the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "The role of the person who created the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_archived",
				Description: "Whether or not the project is archived.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "updated_at",
				Description: "The date and time when project was updated",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_at",
				Description: "The date and time when project was created",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "project_policy",
				Description: "Associated project policy with the project",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getProjectPolicy,
				Transform:   transform.FromValue(),
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

func listProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_project.listProjects", "connection_error", err)
		return nil, err
	}

	project, err := client.Projects()
	if err != nil {
		logger.Error("env0_project.listProjects", "api_error", err)
		return nil, err
	}

	for _, item := range project {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTION

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectId := d.EqualsQualString("id")

	if projectId == "" {
		return nil, nil
	}

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_project.getProject", "connection_error", err)
		return nil, err
	}

	project, err := client.Project(projectId)
	if err != nil {
		plugin.Logger(ctx).Error("env0_project.getProject", "api_error", err)
		return nil, err
	}

	return project, nil
}

func getProjectPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectId := h.Item.(client.Project).Id

	if projectId == "" {
		return nil, nil
	}

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_project.getProjectPolicy", "connection_error", err)
		return nil, err
	}

	result, err := client.Policy(projectId)
	if err != nil {
		plugin.Logger(ctx).Error("env0_project.getProjectPolicy", "api_error", err)
		return nil, err
	}

	return result, nil
}
