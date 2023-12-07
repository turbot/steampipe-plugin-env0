---
title: "Steampipe Table: env0_user - Query env0 Users using SQL"
description: "Allows users to query env0 Users, specifically the user's details, roles, and permissions, providing insights into user management and access control."
---

# Table: env0_user - Query env0 Users using SQL

Env0 is a self-service cloud management platform that enables developers and DevOps teams to manage, govern, and automate their cloud environments. It provides a centralized way to manage users, their roles, and permissions across different cloud environments. Env0 helps you stay informed about the access control and user management of your cloud resources.

## Table Usage Guide

The `env0_user` table provides insights into users within env0's cloud management platform. As a DevOps engineer, explore user-specific details through this table, including user roles, permissions, and associated metadata. Utilize it to uncover information about user management, such as those with administrator roles, the permissions assigned to users, and the verification of user roles.

## Examples

### Basic info
Explore the user details in Env0 to gain insights into their identities, roles, and account creation dates. This can be useful for account management, user segmentation, and understanding user demographics.

```sql+postgres
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

```sql+sqlite
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
Discover the segments of users who have recently joined your platform. This can help in understanding user growth trends and targeting new user onboarding strategies.

```sql+postgres
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

```sql+sqlite
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
  created_at >= datetime('now','-30 day');
```

### List users who haven't logged in to env0 platform in the last 30 days
Determine the users who have been inactive on the env0 platform for the past month. This can help in identifying dormant accounts for potential follow-ups or account management actions.

```sql+postgres
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

```sql+sqlite
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
  last_login <= datetime('now','-30 day');
```

### List users without administrative privileges
Discover the segments that consist of users who do not have administrative privileges. This can be useful in assessing user permissions, ensuring security protocols, and managing user roles within your organization.

```sql+postgres
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

```sql+sqlite
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
Identify instances where users are not currently active. This can be useful for assessing user engagement and identifying potential areas for outreach or account management.

```sql+postgres
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

```sql+sqlite
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