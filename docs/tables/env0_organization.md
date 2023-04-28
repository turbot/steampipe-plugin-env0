# Table: env0_organization

This table returns information about the env0 Organization.

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

### List Organizations created in the last 30 days

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

### List Organizations that have not been updated in the last 30 days

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

### List Organizations with OIDC disabled

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
