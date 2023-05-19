# Table: env0_organization

An env0 organization refers to a logical entity within the env0 platform that serves as a container for managing and organizing resources, environments, and users. Organizations in env0 provide a way to structure and segregate different teams, projects, or departments within a larger context. This table returns information about the env0 Organization.

## Examples

### Basic info

```sql
select
  name,
  id,
  created_by,
  created_at,
  description,
  role,
  photo_url
from
  env0_organization;
```

### List organizations created in the last 30 days

```sql
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  created_at >= now() - interval '30' day;
```

### List organizations that have not been updated in the last 30 days

```sql
select
  name,
  description,
  id,
  created_by,
  created_at,
  updated_at,
  role
from
  env0_organization
where
  updated_at <= now() - interval '30' day;
```

### List organizations with OIDC disabled

```sql
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  not enable_oidc;
```
