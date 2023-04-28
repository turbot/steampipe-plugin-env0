# Table: env0_api_key

This table returns information about the API keys in your env0 Organization.

## Examples

### Basic info

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  organization_role,
  created_at,
  created_by
from
  env0_api_key;
```

### List API keys created in the last 30 days

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  organization_role,
  created_at,
  created_by
from
  env0_api_key
where
  created_at >= now() - interval '30' day;
```

### List API keys which haven't been used in the last 30 days

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  organization_role,
  created_at,
  created_by
from
  env0_api_key
where
  last_used_at <= now() - interval '30' day;
```

### List unused API keys

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  organization_role,
  created_at,
  created_by
from
  env0_api_key
where
  last_used_at is null;
```

### List API keys that have admin privileges

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  organization_role,
  created_at,
  created_by
from
  env0_api_key
where
  organization_role = 'Admin';
```

### List the creator details of each key

```sql
select
  name,
  id,
  api_key_id,
  organization_id,
  created_at,
  created_by
  created_by_user ->> 'name' as created_by_user_name,
  created_by_user ->> 'email' as created_by_user_email
from
  env0_api_key;
```