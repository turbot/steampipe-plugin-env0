# Table: env0_notification

An env0 notification is a mechanism in the env0 platform that enables users to receive real-time updates and alerts regarding various events and activities related to their environments and deployments. Notifications can be configured to notify users about events. This table returns information about the env0 Notification.

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
  env0_notification
where
  type = 'Slack';
```

### List the creator details of a particular notification

```sql
select
  name,
  id,
  created_by_user ->> 'name' as user_name,
  created_by_user ->> 'user_id' as user_id,
  created_by_user ->> 'email' as user_email,
  organization_id,
  created_by,
  type,
  value
from
  env0_notification;
```