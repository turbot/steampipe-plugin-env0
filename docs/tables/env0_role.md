# Table: env0_role

This table returns information about the roles in an env0 Organization. This includes both the default and the custom roles.

## Examples

### Basic info

```sql
select
  name,
  id,
  organization_id,
  is_default_role,
  title
from
  env0_role;
```

### List custom roles

```sql
select
  name,
  id,
  organization_id,
  is_default_role,
  title
from
  env0_role
where
  not is_default_role;
```

### List the permissions associated with a role

```sql
select
  name,
  id,
  organization_id,
  is_default_role,
  jsonb_pretty(permissions)
from
  env0_role;
```

### List the custom roles that provide organizational view access

```sql
select
  name,
  id,
  organization_id,
  is_default_role,
  jsonb_pretty(permissions)
from
  env0_role
where
  not is_default_role
  and
  (
    permissions
  )
  @ > '"VIEW_ORGANIZATION"';
```

