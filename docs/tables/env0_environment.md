# Table: env0_environment

An env0 environment is a virtual representation of a software development or deployment environment. It is created within the env0 platform to facilitate the management and orchestration of infrastructure resources, configurations, and deployments. An environment in env0 typically consists of various components such as cloud providers, infrastructure templates, configurations, and application code. This table returns information about the env0 Environment.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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

```sql
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

### Get the log details of the latest deployment

```sql
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
