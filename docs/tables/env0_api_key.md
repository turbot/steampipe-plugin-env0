---
title: "Steampipe Table: env0_api_key - Query env0 API Keys using SQL"
description: "Allows users to query API Keys in env0, specifically the details of each API key such as its ID, name, role, and status."
---

# Table: env0_api_key - Query env0 API Keys using SQL

The env0 API Key is a resource in env0 that allows users to authenticate and authorize API calls. This key is associated with a specific role and can be used to perform actions based on the permissions granted to that role. The status of the API key can be active or inactive, based on which it can be used to make API calls.

## Table Usage Guide

The `env0_api_key` table provides insights into API keys within env0. As a DevOps engineer, explore API key-specific details through this table, including their ID, name, role, and status. Utilize it to uncover information about API keys, such as those associated with specific roles, and the status of these keys.

## Examples

### Basic info
Explore which API keys are linked to your organization, when they were created, and by whom. This can help you manage access and roles within your organization effectively.

```sql+postgres
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

```sql+sqlite
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
Discover the recent API keys that have been created within the last month. This can be useful for auditing purposes, to track the creation of new keys and maintain the security of your environment.

```sql+postgres
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

```sql+sqlite
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
  created_at >= datetime('now', '-30 day');
```

### List API keys which haven't been used in the last 30 days
Discover the API keys that have remained inactive for the past 30 days. This information could be useful for identifying potential security risks or unused resources in your organization.

```sql+postgres
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

```sql+sqlite
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
  last_used_at <= datetime('now', '-30 day');
```

### List unused API keys
Discover the segments that have unused API keys, enabling you to manage and clean up your API key inventory efficiently. This is particularly useful in maintaining security and reducing unnecessary clutter in your system.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have administrative privileges by analyzing API keys. This could be useful in managing access and maintaining security within your organization.

```sql+postgres
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

```sql+sqlite
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
Explore the individuals behind each key creation to gain insights into their associated organizations and the timing of their activities. This can be instrumental in tracking key usage, understanding organizational relationships, and ensuring accountability.

```sql+postgres
select
  name,
  id,
  api_key_id,
  organization_id,
  created_at,
  created_by,
  created_by_user ->> 'name' as created_by_user_name,
  created_by_user ->> 'email' as created_by_user_email
from
  env0_api_key;
```

```sql+sqlite
select
  name,
  id,
  api_key_id,
  organization_id,
  created_at,
  created_by,
  json_extract(created_by_user, '$.name') as created_by_user_name,
  json_extract(created_by_user, '$.email') as created_by_user_email
from
  env0_api_key;
```