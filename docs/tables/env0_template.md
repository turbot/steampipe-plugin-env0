# Table: env0_template

This table returns information about the Env0 Template.

## Examples

### Basic info

```sql
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template;
```

### List templates created in the last 30 days

```sql
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  created_at >= now() - interval '30' day;
```

### List templates that have not been updated in the last 30 days

```sql
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  updates_at >= now() - interval '30' day;
```
