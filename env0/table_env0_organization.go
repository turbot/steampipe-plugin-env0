package env0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEnv0Organization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "env0_organization",
		Description: "Returns information about the Env0 organization.",
		List: &plugin.ListConfig{
			Hydrate: listOrganizations,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "Organization ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_ttl",
				Description: "Max TTL of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default_ttl",
				Description: "Default TTL of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "do_not_report_skipped_status_checks",
				Description: "Check if the organization does not report skipped status checks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "do_not_consider_merge_commits_for_pr_plans",
				Description: "Check if the organization does not consider merge commits for PR plans.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_oidc",
				Description: "Check whether OIDC is enabled for the organziation.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enforce_pr_commenter_permissions",
				Description: "Check whether the organization has enforced PR commenter permissions.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "Organization description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "photo_url",
				Description: "The URL of the organization's display photo.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "The creator of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Date and time when the organziation was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_at",
				Description: "Date and time when the organziation last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "role",
				Description: "Organization role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_self_hosted_k8s",
				Description: "Returns true if the orgization is a self hosted k8.",
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

func listOrganizations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("env0_organization.listOrganizations", "connection_error", err)
		return nil, err
	}

	organization, err := client.Organization()
	if err != nil {
		logger.Error("env0_organization.listOrganizations", "api_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, organization)
	if d.RowsRemaining(ctx) == 0 {
		return nil, nil
	}

	return nil, nil
}
