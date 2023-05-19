# Table: env0_team

In env0, a team refers to a group of users who collaborate and work together within an organization. Teams in env0 help organize users and facilitate collaboration on projects, environments, and deployments. This table returns information about the teams in an env0 Organization.

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

### Get the last login details of the users in a team

```sql
select
  name,
  id,
  description,
  u ->> 'name' as user_name,
  u ->> 'email' as user_email,
  u ->> 'user_id' as user_id,
  u ->> 'LastLogin' as user_last_login,
from
  env0_team,
  jsonb_array_elements(users) as u;
```
