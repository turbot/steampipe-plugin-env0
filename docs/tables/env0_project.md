---
title: "Steampipe Table: env0_project - Query Env0 Projects using SQL"
description: "Allows users to query Env0 Projects, specifically to retrieve the details of each project, providing insights into project settings, users, and environments."
---

# Table: env0_project - Query Env0 Projects using SQL

Env0 is a platform that enables Infrastructure as Code (IaC) as part of the software development lifecycle. It allows developers to manage and govern their cloud resources, providing an automated, collaborative, and programmatically controlled cloud management solution. Env0 supports a range of IaC frameworks, including Terraform, and offers features such as cost estimation and policy enforcement.

## Table Usage Guide

The `env0_project` table provides insights into projects within the Env0 platform. As a DevOps engineer, you can explore project-specific details through this table, including project settings, associated users, and environments. Utilize it to uncover information about projects, such as those with specific environment settings, the users associated with each project, and the status of each project.

## Examples

### Basic info
Explore which projects have been archived within your organization. This allows for a quick overview of project status and can help identify inactive projects for potential cleanup or resource reallocation.

```sql+postgres
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project;
```

```sql+sqlite
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project;
```

### List project created in the last 30 days
Discover projects that have been initiated within the past month. This allows you to stay updated with new developments and assess the latest projects in your organization.

```sql+postgres
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project
where
  created_at >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project
where
  created_at >= datetime('now', '-30 day');
```

### List projects that have not been updated in the last 30 days
Explore projects that have been inactive for the past 30 days. This can be particularly useful for identifying and managing outdated or neglected projects, thereby improving overall operational efficiency.

```sql+postgres
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  updated_at <= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  updated_at <= datetime('now', '-30 day');
```

### List unarchived projects
Discover the segments that are currently active in your project portfolio by identifying projects that are not archived. This aids in focusing resources and attention on ongoing initiatives, ensuring efficient project management.

```sql+postgres
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  is_archived is false;
```

```sql+sqlite
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  is_archived = 0;
```

### List projects that do not implement continuous deployment
Determine the areas in which continuous deployment is not implemented for projects. This analysis can help in identifying potential gaps in the development process and drive improvements in the deployment strategy.

```sql+postgres
select
  created_by,
  created_by_user,
  id,
  is_archived,
  name,
  organization_id
  role
from
  env0_project
where
  project_policy ->> 'continuousDeploymentDefault' = 'false';
```

```sql+sqlite
select
  created_by,
  created_by_user,
  id,
  is_archived,
  name,
  organization_id,
  role
from
  env0_project
where
  json_extract(project_policy, '$.continuousDeploymentDefault') = 'false';
```

### Get all project policy settings of specific project
Identify and understand the policy settings for a specific project, including the maximum number of environments and default approval requirements. This can be useful for project managers to monitor and control project settings, ensuring efficient resource utilization and adherence to project guidelines.

```sql+postgres
select
  name,
  organization_id,
  project_policy ->> 'numberOfEnvironmentsTotal' as max_no_of_environments,
  project_policy ->> 'requiresApprovalDefault' as requires_approval_default,
  project_policy ->> 'numberOfEnvironments' as no_of_environment,
  project_policy ->> 'continuousDeploymentDefault' as continuous_deployment_default,
  project_policy ->> 'maxTtl' as max_ttl,
  project_policy ->> 'defaultTtl' as default_ttl
from
  env0_project
where
  id = '4a639364-1234-4eee-5678-ad38e1c1ccee';
```

```sql+sqlite
select
  name,
  organization_id,
  json_extract(project_policy, '$.numberOfEnvironmentsTotal') as max_no_of_environments,
  json_extract(project_policy, '$.requiresApprovalDefault') as requires_approval_default,
  json_extract(project_policy, '$.numberOfEnvironments') as no_of_environment,
  json_extract(project_policy, '$.continuousDeploymentDefault') as continuous_deployment_default,
  json_extract(project_policy, '$.maxTtl') as max_ttl,
  json_extract(project_policy, '$.defaultTtl') as default_ttl
from
  env0_project
where
  id = '4a639364-1234-4eee-5678-ad38e1c1ccee';
```