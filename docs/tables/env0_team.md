---
title: "Steampipe Table: env0_team - Query env0 Teams using SQL"
description: "Allows users to query Teams in env0, specifically the team name, id, and other related details, providing insights into team management and access control."
---

# Table: env0_team - Query env0 Teams using SQL

env0 is a cloud management platform that provides Infrastructure as Code (IaC) automation, governance, and collaboration for cloud environments. It enables teams to manage and control cloud resources across multiple environments. Teams in env0 represent a group of users with specific access permissions to the platform's resources.

## Table Usage Guide

The `env0_team` table provides insights into Teams within env0. As a cloud manager or DevOps engineer, explore team-specific details through this table, including team id, name, and associated user count. Utilize it to understand team structure, manage access control, and ensure appropriate resource allocation among different teams.

## Examples

### Basic info
Explore your team's basic information, such as names and associated IDs, to understand the structure and organization of your team within the Env0 platform. This can help in managing team resources and planning for future team expansions.

```sql+postgres
select
  name,
  id,
  description,
  organization_id
from
  env0_team;
```

```sql+sqlite
select
  name,
  id,
  description,
  organization_id
from
  env0_team;
```

### List the details of the users in a team
Explore the details of team members, such as their names and email addresses, to gain a comprehensive overview of your team's composition. This can be particularly useful for team management and communication purposes.

```sql+postgres
select
  env0_team.name,
  env0_team.id,
  env0_team.description,
  u ->> 'name' as user_name,
  u ->> 'email' as user_email,
  u ->> 'user_id' as user_id
from
  env0_team,
  jsonb_array_elements(users) as u;
```

```sql+sqlite
select
  env0_team.name,
  env0_team.id,
  env0_team.description,
  json_extract(u.value, '$.name') as user_name,
  json_extract(u.value, '$.email') as user_email,
  json_extract(u.value, '$.user_id') as user_id
from
  env0_team,
  json_each(users) as u;
```

### Get the last login details of the users in a team
Identify instances where you need to track user activity in a team by checking their last login details. This can be useful for auditing purposes or to monitor user engagement within the team.

```sql+postgres
select
  env0_team.name,
  env0_team.id,
  env0_team.description,
  u ->> 'name' as user_name,
  u ->> 'email' as user_email,
  u ->> 'user_id' as user_id,
  u ->> 'LastLogin' as user_last_login,
from
  env0_team,
  jsonb_array_elements(users) as u;
```

```sql+sqlite
select
  env0_team.name,
  env0_team.id,
  env0_team.description,
  json_extract(u.value, '$.name') as user_name,
  json_extract(u.value, '$.email') as user_email,
  json_extract(u.value, '$.user_id') as user_id,
  json_extract(u.value, '$.LastLogin') as user_last_login
from
  env0_team,
  json_each(users) as u;
```