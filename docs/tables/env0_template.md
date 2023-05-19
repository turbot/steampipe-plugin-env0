# Table: env0_template

Env0 templates provide a declarative approach to defining infrastructure and environment configurations. They describe the desired resources, services, and settings needed for an environment to be provisioned and deployed. This table returns information about the env0 templates.

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
  updated_at >= now() - interval '30' day;
```

### List templates that uses GitHub or GitLab enterprise

```sql
select
  name,
  id,
  author,
  created_at,
  organization_id,
  project_ids,
  terraform_version
from
  env0_template
where
  is_github_enterprise or is_gitlab_enterprise;
```
