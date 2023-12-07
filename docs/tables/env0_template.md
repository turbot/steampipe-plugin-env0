---
title: "Steampipe Table: env0_template - Query env0 Templates using SQL"
description: "Allows users to query env0 Templates, providing insights into the infrastructure as code (IaC) templates used in the env0 platform."
---

# Table: env0_template - Query env0 Templates using SQL

An env0 Template is a predefined infrastructure as code (IaC) template that can be used to create and manage resources on the env0 platform. These templates are typically written in Terraform or Terragrunt, and define the infrastructure components needed for an application. The templates are reusable, version-controlled, and can be parameterized to create unique environments based on the same template.

## Table Usage Guide

The `env0_template` table provides insights into the infrastructure as code (IaC) templates used within the env0 platform. As a DevOps engineer, you can explore template-specific details through this table, including the template's name, version, and associated parameters. Utilize it to manage and understand the infrastructure components defined in each template, allowing for better control and visibility over your env0 environments.

## Examples

### Basic info
Explore which templates are created by whom and when, along with their associated organizations and projects. This could be useful to understand how your infrastructure as code templates are being used across different enterprise version control systems.

```sql+postgres
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template;
```

```sql+sqlite
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template;
```

### List templates created in the last 30 days
Discover the segments that have been recently added by identifying templates created within the past month. This can help keep track of new additions and understand recent changes in your environment.

```sql+postgres
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  created_at >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  created_at >= datetime('now', '-30 day');
```

### List templates that have not been updated in the last 30 days
Assess the elements within your environment to identify any templates that haven't been updated in the past month. This can be useful to ensure all templates are up-to-date and adhere to current standards or practices.

```sql+postgres
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  updated_at >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  id,
  author,
  created_at,
  organization_id,
  is_github_enterprise,
  is_gitlab_enterprise,
  project_ids,
  terraform_version
from
  env0_template
where
  updated_at >= datetime('now', '-30 day');
```

### List templates that uses GitHub or GitLab enterprise
Explore the templates created for organizations that utilize either GitHub or GitLab enterprise versions. This is particularly useful for identifying and analyzing the use of these enterprise platforms in your projects.

```sql+postgres
select
  name,
  id,
  author,
  created_at,
  organization_id,
  project_ids,
  terraform_version
from
  env0_template
where
  is_github_enterprise or is_gitlab_enterprise;
```

```sql+sqlite
select
  name,
  id,
  author,
  created_at,
  organization_id,
  project_ids,
  terraform_version
from
  env0_template
where
  is_github_enterprise = 1 or is_gitlab_enterprise = 1;
```