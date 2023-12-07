---
title: "Steampipe Table: env0_environment - Query Env0 Environments using SQL"
description: "Allows users to query Env0 Environments, specifically to retrieve detailed information about the environments, including their ID, name, project ID, and other related data."
---

# Table: env0_environment - Query Env0 Environments using SQL

Env0 is a self-service cloud management platform that provides infrastructure as code (IaC) environments on demand. It allows teams to manage and provision their own environments, while maintaining control and governance. Env0 supports multiple IaC frameworks and cloud providers, enabling teams to use their preferred technologies and processes.

## Table Usage Guide

The `env0_environment` table provides insights into Environments within Env0. As a DevOps engineer, you can explore environment-specific details through this table, including the environment's ID, name, project ID, and other related data. Use it to uncover information about environments, such as their current status, the associated project, and the last time they were updated.

## Examples

### Basic info
Explore which environments require approval for changes and identify instances where continuous deployment is enabled. This can aid in understanding the management and control mechanisms in place for your projects.

```sql+postgres
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  is_archived
from
  env0_environment;
```

```sql+sqlite
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  is_archived
from
  env0_environment;
```

### List environments which are not in archived state
Discover the environments that are currently active and not archived. This can be useful for maintaining an overview of all ongoing projects and ensuring that resources are not being wasted on inactive environments.

```sql+postgres
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  not is_archived;
```

```sql+sqlite
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  not is_archived;
```

### List inactive environments
Explore which environments are inactive in your project, helping you identify areas for clean-up or potential resource optimization. This is useful in managing resources and maintaining an efficient, organized environment structure.

```sql+postgres
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  status <> 'Active';
```

```sql+sqlite
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  status <> 'Active';
```

### List environments that bypass approval requirement from an authorized user
Explore which environments are set to bypass the approval requirement, allowing for a more streamlined deployment process and potentially reducing bottlenecks. This can be beneficial in scenarios where rapid deployment is critical.

```sql+postgres
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  not requires_approval;
```

```sql+sqlite
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  not requires_approval;
```

### List environments that do not implement continuous deployment
Determine the areas in your project where continuous deployment is not implemented. This can help identify potential bottlenecks in your development process and areas for improvement.

```sql+postgres
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  not continuous_deployment;
```

```sql+sqlite
select
  name,
  id,
  project_id,
  continuous_deployment,
  pull_request_plan_deployments,
  requires_approval,
  status,
  workspace_name,
  is_archived
from
  env0_environment
where
  continuous_deployment = 0;
```

### Get the log details of the latest deployment
Explore the latest deployment's log details to identify if there were any issues or irregularities. This is particularly useful in situations where continuous deployment is not used, allowing for more thorough tracking and troubleshooting of each deployment.

```sql+postgres
select
  name,
  id,
  project_id,
  workspace_name,
  is_archived,
  latest_deployment_log ->> 'Id' as latest_deployment_log_id,
  latest_deployment_log ->> 'BlueprintId' as latest_deployment_log_blueprint_id,
  latest_deployment_log ->> 'Type' as latest_deployment_log_type
from
  env0_environment
where
  not continuous_deployment;
```

```sql+sqlite
select
  name,
  id,
  project_id,
  workspace_name,
  is_archived,
  json_extract(latest_deployment_log, '$.Id') as latest_deployment_log_id,
  json_extract(latest_deployment_log, '$.BlueprintId') as latest_deployment_log_blueprint_id,
  json_extract(latest_deployment_log, '$.Type') as latest_deployment_log_type
from
  env0_environment
where
  not continuous_deployment;
```