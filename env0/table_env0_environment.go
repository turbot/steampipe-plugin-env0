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
		Description: "Returns information about the Env0 environment.",
		List: &plugin.ListConfig{
			Hydrate: listEnvironments,
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("id"),
		// 	Hydrate:    getTemplate,
		// },
		Columns: []*plugin.Column{
			{
				Name:        "data",
				Description: "Name of the template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
			},
			// {
			// 	Name:        "id",
			// 	Description: "Unique identifier of the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "organization_id",
			// 	Description: "Unique identifier of the organization.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "author",
			// 	Description: "Name of the author of the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "author_id",
			// 	Description: "Unique identifier of the author.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "created_at",
			// 	Description: "The date and time when template was created",
			// 	Type:        proto.ColumnType_TIMESTAMP,
			// },
			// {
			// 	Name:        "href",
			// 	Description: "Href link to the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "description",
			// 	Description: "Description associated with the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },

			// {
			// 	Name:        "path",
			// 	Description: "The Path of the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "revision",
			// 	Description: "Revision of the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "project_id",
			// 	Description: "The associated project id with the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "project_ids",
			// 	Description: "The associated projects id with the template.",
			// 	Type:        proto.ColumnType_JSON,
			// },
			// {
			// 	Name:        "repository",
			// 	Description: "The associated version control repository with the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "retry",
			// 	Description: "Retry associated with the template.",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   transform.FromField("Retry"),
			// },
			// {
			// 	Name:        "ssh_keys",
			// 	Description: "Associated ssh keys with the template.",
			// 	Type:        proto.ColumnType_JSON,
			// },
			// {
			// 	Name:        "type",
			// 	Description: "Type of the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "github_installation_id",
			// 	Description: "The GitHub installation of the template.",
			// 	Type:        proto.ColumnType_INT,
			// },
			// {
			// 	Name:        "is_gitlab_enterprise",
			// 	Description: "Whether or not template belongs to GitLab enterprise.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "is_github_enterprise",
			// 	Description: "Whether or not template belongs to GitHub enterprise.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "token_id",
			// 	Description: "Token ID associated with the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "updated_at",
			// 	Description: "The date and time when template was updated",
			// 	Type:        proto.ColumnType_TIMESTAMP,
			// },
			// {
			// 	Name:        "terraform_version",
			// 	Description: "Terraform version associated with the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "terragrunt_version",
			// 	Description: "Specifies the version of terragrunt that should be used when running a terraform deployment by the template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "is_terragrunt_run_all",
			// 	Description: "Whether or not terragrunt will execute all the modules associated with the template.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "is_deleted",
			// 	Description: "Whether or not the template is deleted.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "is_azure_devops",
			// 	Description: "Whether the template is associated with Azure DevOps deployment.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "bitbucket_client_Key",
			// 	Description: "Bitbucket client key used by template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "is_bitbucket_server",
			// 	Description: "Whether or not template is using Bitbucket server.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "file_name",
			// 	Description: "File name associated with template.",
			// 	Type:        proto.ColumnType_STRING,
			// },
		},
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

	result, err := client.Environments()
	if err != nil {
		logger.Error("env0_environment.listEnvironments", "api_error", err)
		return nil, err
	}

	for _, item := range result {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTION

// func getTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	logger := plugin.Logger(ctx)
// 	templateId := d.EqualsQualString("id")

// 	if templateId == "" {
// 		return nil, nil
// 	}

// 	// Create client
// 	client, err := connect(ctx, d)
// 	if err != nil {
// 		logger.Error("env0_environment.getTemplate", "connection_error", err)
// 		return nil, err
// 	}

// 	result, err := client.Template(templateId)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("env0_environment.getTemplate", "api_error", err)
// 		return nil, err
// 	}

// 	return result, nil
// }
