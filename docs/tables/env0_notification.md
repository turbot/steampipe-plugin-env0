---
title: "Steampipe Table: env0_notification - Query Env0 Notifications using SQL"
description: "Allows users to query Notifications in Env0, specifically the notification settings for different environments and templates, providing insights into communication preferences and potential anomalies."
---

# Table: env0_notification - Query Env0 Notifications using SQL

Env0 Notifications is a feature within the Env0 platform that allows users to manage and customize their communication preferences for different environments and templates. It provides a centralized way to set up and manage alerts for various Env0 resources, including environments, projects, and templates. Env0 Notifications helps users stay informed about the status and updates of their Env0 resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `env0_notification` table provides insights into notification settings within the Env0 platform. As a DevOps engineer, explore notification-specific details through this table, including alert preferences, associated environments, and templates. Utilize it to uncover information about notifications, such as those related to specific environments, the communication preferences for different templates, and verification of alert conditions.

## Examples

### Basic info
Explore the notifications within your environment to understand their origin and type. This can help you assess who or what is generating these notifications, and the kind of information they contain, which is beneficial for managing and monitoring your environment.

```sql+postgres
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

```sql+sqlite
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
Explore which notifications are created by users in your organization on Slack. This can help you assess the communication flow and understand the context of each notification.

```sql+postgres
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

```sql+sqlite
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
Explore who created a specific notification. This can aid in tracing the origin of certain alerts and assist in communication within the team.

```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  json_extract(created_by_user, '$.name') as user_name,
  json_extract(created_by_user, '$.user_id') as user_id,
  json_extract(created_by_user, '$.email') as user_email,
  organization_id,
  created_by,
  type,
  value
from
  env0_notification;
```