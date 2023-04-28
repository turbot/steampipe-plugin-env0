# Table: env0_project

This table returns information about the env0 Organization.

## Examples

### Basic info

```sql
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

```sql
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

### List projects that have not been updated in the last 30 days

```sql
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

### List all projects archive is set to false'

```sql
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
  is_archived is false
```

### List all projects where continuous deployment default is set to false

```sql
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
  project_policy ->> 'continuousDeploymentDefault' = 'false'
```

### Get all project policy settings of specific project

```sql
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
  id = '4a639364-1234-4eee-5678-ad38e1c1ccee'
```
