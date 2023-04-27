# Table: env0_project

This table returns information about the Env0 Organization.

## Examples

### Basic info

```sql
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project;
```

### List project created in the last 30 days

```sql
select
  name,
  id,
  created_by,
  created_by_user,
  organization_id,
  is_archived
from
  env0_project
where
  created_at >= now() - interval '30' day;
```

### List projects that have not been updated in the last 30 days

```sql
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  updated_at <= now() - interval '30' day;
```

### List all projects archive is set to false'

```sql
select
  name,
  id,
  created_by,
  created_by_user,
  role,
  organization_id,
  is_archived
from
  env0_project
where
  is_archived is false
```
