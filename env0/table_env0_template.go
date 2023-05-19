package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableenv0Template(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_template",
		Description: "Returns information about the env0 template.",
		List: &plugin.ListConfig{
			Hydrate: listTemplates,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTemplate,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "Unique identifier of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "The organization ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "author",
				Description: "Name of the author of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "author_id",
				Description: "Unique identifier of the author.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date and time when template was created",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "href",
				Description: "Href link to the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description associated with the template.",
				Type:        proto.ColumnType_STRING,
			},

			{
				Name:        "path",
				Description: "The Path of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "revision",
				Description: "Revision of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "The associated project id with the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_ids",
				Description: "The associated projects id with the template.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "repository",
				Description: "The associated version control repository with the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "retry",
				Description: "Retry associated with the template.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ssh_keys",
				Description: "The associated ssh keys with the template.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "type",
				Description: "Type of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "github_installation_id",
				Description: "The GitHub installation of the template.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "is_gitlab_enterprise",
				Description: "Specifies whether or not template belongs to GitLab enterprise.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_github_enterprise",
				Description: "Specifies whether or not template belongs to GitHub enterprise.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "token_id",
				Description: "Token ID associated with the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The date and time when template was updated",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "terraform_version",
				Description: "Terraform version associated with the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "terragrunt_version",
				Description: "Specifies the version of terragrunt that should be used when running a terraform deployment by the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_terragrunt_run_all",
				Description: "Specifies whether or not terragrunt will execute all the modules associated with the template.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_deleted",
				Description: "Specifies whether or not the template is deleted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_azure_devops",
				Description: "Specifies whether the template is associated with Azure DevOps deployment.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "bitbucket_client_Key",
				Description: "The Bitbucket client key used by template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_bitbucket_server",
				Description: "Specifies whether or not template is using Bitbucket server.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "file_name",
				Description: "File name associated with template.",
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

func listTemplates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_template.listTemplates", "connection_error", err)
		return nil, err
	}

	result, err := client.Templates()
	if err != nil {
		logger.Error("env0_template.listTemplates", "api_error", err)
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

func getTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	templateId := d.EqualsQualString("id")

	if templateId == "" {
		return nil, nil
	}

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_template.getTemplate", "connection_error", err)
		return nil, err
	}

	result, err := client.Template(templateId)
	if err != nil {
		plugin.Logger(ctx).Error("env0_template.getTemplate", "api_error", err)
		return nil, err
	}

	return result, nil
}
