---
title: "Steampipe Table: env0_role - Query Env0 Roles using SQL"
description: "Allows users to query Roles in Env0, specifically the role names, permissions, and associated users, providing insights into role-based access control within the Env0 environment."
---

# Table: env0_role - Query Env0 Roles using SQL

Env0 is a platform that manages and orchestrates infrastructure as code (IaC) environments. It allows developers to plan, build, and deploy cloud resources using IaC frameworks. Env0 simplifies the management of environments, offering features like role-based access control, cost management, and policy governance.

## Table Usage Guide

The `env0_role` table provides insights into roles within the Env0 platform. As a DevOps engineer, explore role-specific details through this table, including role names, permissions, and associated users. Utilize it to manage and monitor role-based access control within your Env0 environments.

## Examples

### Basic info
Analyze the settings to understand the roles within your Env0 environment, including their titles and whether they are default roles, to effectively manage access and permissions.

```sql+postgres
select
  name,
  id,
  organization_id,
  is_default_role,
  title
from
  env0_role;
```

```sql+sqlite
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
Discover the segments that consist of custom roles within your organization. This is useful to understand specific role-based access controls, particularly those that are not default, thereby enhancing your organization's security protocol.

```sql+postgres
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

```sql+sqlite
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
Explore which responsibilities are attached to a particular role in an organization. This is essential for managing user access and ensuring appropriate security measures are in place.

```sql+postgres
select
  name,
  id,
  organization_id,
  is_default_role,
  jsonb_pretty(permissions)
from
  env0_role;
```

```sql+sqlite
select
  name,
  id,
  organization_id,
  is_default_role,
  permissions
from
  env0_role;
```

### List the custom roles that provide organizational view access
Discover the custom roles that have been granted access to view organizational details. This is useful for understanding how permissions are distributed within your organization and identifying any potential security risks.

```sql+postgres
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

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```

### List the custom roles that provide access to edit organization settings
Determine the custom roles that have the ability to modify organization settings. This is useful for managing access control and ensuring only authorized roles have the ability to make changes to the organization's configurations.

```sql+postgres
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

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```