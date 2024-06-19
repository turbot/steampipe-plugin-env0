package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableenv0Environment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_environment",
		Description: "Returns information about the env0 environment.",
		List: &plugin.ListConfig{
			Hydrate: listEnvironments,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEnvironment,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The number uniquely identifying the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "A unique ID of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "workspace_name",
				Description: "The name of the workspace associated to the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "requires_approval",
				Description: "A boolean value indicating whether changes to the environment require approval before they can be applied.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "continuous_deployment",
				Description: "Specifies whether or not continuous deployment value is set in the environment.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "pull_request_plan_deployments",
				Description: "Specifies whether automatic deployments of infrastructure change is enabled when a pull request is created or updated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "auto_deploy_on_path_changes_only",
				Description: "Specifies whether or not automatic deployments of infrastructure changes should be triggered only when changes are made to specific files or directories.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "auto_deploy_by_custom_glob",
				Description: "Specifies a custom glob pattern that determines which files should trigger an automatic deployment when changes are detected.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "lifespan_end_at",
				Description: "Represents the end time of an environment's lifespan",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LifespanEndAt").NullIfZero(),
			},
			{
				Name:        "latest_deployment_log_id",
				Description: "Represents the ID of the most recent deployment log.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "latest_deployment_log",
				Description: "The most recent deployment log for a specific environment or deployment.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "terragrunt_working_directory",
				Description: "Specifies the directory where Terragrunt will look for the terragrunt.hcl configuration file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vcs_commands_alias",
				Description: "Specifies alias for the VCS commands used for deploying infrastructure changes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "blueprint_id",
				Description: "Specifies the ID of the blueprint associated with the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_remote_backend",
				Description: "Specifies whether the Terraform backend configuration uses a remote backend or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_archived",
				Description: "Specifies whether an environment is set to archived or not.",
				Type:        proto.ColumnType_BOOL,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listEnvironments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_environment.listEnvironments", "connection_error", err)
		return nil, err
	}

	environment, err := client.Environments()
	if err != nil {
		logger.Error("env0_environment.listEnvironments", "api_error", err)
		return nil, err
	}

	for _, item := range environment {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTION

func getEnvironment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	environmentId := d.EqualsQualString("id")

	if environmentId == "" {
		return nil, nil
	}

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_environment.getEnvironment", "connection_error", err)
		return nil, err
	}

	environment, err := client.Environment(environmentId)
	if err != nil {
		plugin.Logger(ctx).Error("env0_environment.getEnvironment", "api_error", err)
		return nil, err
	}

	return environment, nil
}
