# Table: env0_environment

This table returns information about the env0 Environment.

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