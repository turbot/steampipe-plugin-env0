# Table: env0_team

This table returns information about the teams in an Env0 Organization.

## Examples

### Basic info

```sql
select
  name,
  id,
  description,
  organization_id
from
  env0_team;
```

### List the details of the users in a team

```sql
select
  name,
  id,
  description,
  u ->> 'name' as user_name,
  u ->> 'email' as user_email,
  u ->> 'user_id' as user_id
from
  env0_team,
  jsonb_array_elements(users) as u;
```