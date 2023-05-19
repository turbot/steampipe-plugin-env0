# Table: env0_user

In env0, a user refers to an individual who has access to the env0 platform and participates in the management, configuration, and deployment of environments and resources. This table returns information about the users in an env0 Organization.

## Examples

### Basic info

```sql
select
  name,
  email,
  family_name,
  given_name,
  user_id,
  created_at,
  picture,
  role
from
  env0_user;
```

### List users created in the last 30 days

```sql
select
  name,
  email,
  family_name,
  given_name,
  user_id,
  created_at,
  picture,
  role
from
  env0_user
where
  created_at >= now() - interval '30' day;
```

### List users who haven't logged in to env0 platform in the last 30 days

```sql
select
  name,
  email,
  last_login,
  user_id,
  created_at,
  picture,
  role
from
  env0_user
where
  last_login <= now() - interval '30' day;
```

### List users without administrative privileges

```sql
select
  name,
  email,
  last_login,
  user_id,
  created_at,
  picture,
  role
from
  env0_user
where
  role <> 'Admin';
```

### List inactive users

```sql
select
  name,
  email,
  last_login,
  user_id,
  created_at,
  picture,
  role,
  status
from
  env0_user
where
  status <> 'Active';
```
