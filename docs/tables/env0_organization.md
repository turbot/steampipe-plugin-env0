---
title: "Steampipe Table: env0_organization - Query env0 Organizations using SQL"
description: "Allows users to query env0 Organizations, specifically the unique id, name, and role of each organization, providing insights into the organization management in env0."
---

# Table: env0_organization - Query env0 Organizations using SQL

An env0 Organization is a primary entity in env0 which groups a set of users, projects, and templates. It provides a logical separation between different workspaces and is used to manage access control, billing, and settings at an organization level. Each organization has a unique id, a name, and a role associated with it.

## Table Usage Guide

The `env0_organization` table provides insights into Organizations within env0. As a DevOps engineer, explore organization-specific details through this table, including unique id, name, and role. Utilize it to manage and monitor the access control, billing, and settings at an organization level.

## Examples

### Basic info
Discover the segments that were created within your organization, including who created them and when. This can help you gain insights into the structure and activity of your organization.

```sql+postgres
select
  name,
  id,
  created_by,
  created_at,
  description,
  role,
  photo_url
from
  env0_organization;
```

```sql+sqlite
select
  name,
  id,
  created_by,
  created_at,
  description,
  role,
  photo_url
from
  env0_organization;
```

### List organizations created in the last 30 days
Explore which organizations have been established in the recent month. This is useful for keeping track of new additions and ensuring they are properly set up and managed.

```sql+postgres
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  created_at >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  created_at >= datetime('now', '-30 day');
```

### List organizations that have not been updated in the last 30 days
Explore which organizations have remained unchanged in the past month. This is useful for identifying potential areas of stagnation or inactivity within your business.

```sql+postgres
select
  name,
  description,
  id,
  created_by,
  created_at,
  updated_at,
  role
from
  env0_organization
where
  updated_at <= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  description,
  id,
  created_by,
  created_at,
  updated_at,
  role
from
  env0_organization
where
  updated_at <= datetime('now', '-30 day');
```

### List organizations with OIDC disabled
Discover the organizations that have not enabled OpenID Connect (OIDC) for identity authentication. This is useful for identifying potential security risks and ensuring that all organizations are using secure authentication methods.

```sql+postgres
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  not enable_oidc;
```

```sql+sqlite
select
  name,
  description,
  id,
  created_by,
  created_at,
  role
from
  env0_organization
where
  enable_oidc = 0;
```