# Table: env0_notification

This table returns information about the env0 Notification.

## Examples

### Basic info

```sql
select
  name,
  id,
  created_by,
  organization_id,
  created_by_user,
  type,
  value
from
  env0_notification;
```

### List slack notifications

```sql
select
  name,
  id,
  created_by,
  organization_id,
  created_by_user,
  type,
  value
from
  env0_notification;
where
  type = 'Slack';
```
