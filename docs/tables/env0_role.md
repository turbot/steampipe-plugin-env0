# Table: env0_role

In env0, a role refers to a predefined set of permissions and access rights that can be assigned to users within an organization. Roles define what actions a user is allowed to perform within the env0 platform.

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
  @> '"VIEW_ORGANIZATION"';
```

### List the custom roles that provide access to edit organization settings

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
  @> '"EDIT_ORGANIZATION_SETTINGS"';
```