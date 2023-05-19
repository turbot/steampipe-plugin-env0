---
organization: Turbot
category: ["iac"]
icon_url: "/images/plugins/turbot/env0.svg"
brand_color: "#00EDB9"
display_name: "env0"
short_name: "env0"
description: "Steampipe plugin to query projects, teams, users and more from env0."
og_description: "Query env0 with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/env0-social-graphic.png"
---

# env0 + Steampipe

[env0](https://env0.com) is an automation platform for cloud environments based on infrastructure-as-code templates. env0 combines an easy to use interface with powerful governance tools and cost control so that you, or any member of your team, can quickly and easily deploy and manage environments in the cloud.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your env0 organization details:

```sql
 select
  name,
  id,
  created_by,
  created_at,
  description,
  role
from
  env0_organization;
```

```
+--------+--------------------------------------+-------------------------------------+---------------------------+-------------------+-------+
| name   | id                                   | created_by                          | created_at                | description       | role  |
+--------+--------------------------------------+-------------------------------------+---------------------------+-------------------+-------+
| My org | 20a6g055-c90e-3630-9121-a89274f71324 | google-oauth2|116303011440913654789 | 2023-04-26T18:21:33+05:30 | code deployment   | Admin |
+--------+--------------------------------------+-------------------------------------+---------------------------+-------------------+-------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/env0/tables)**

## Quick start

### Install

Download and install the latest env0 plugin:

```sh
steampipe plugin install env0
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                          |
| ----------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | env0 requires an API Key and an API Secret for all requests.                                                                                                                                                                                         |
| Permissions | The organization administrator creates the API keys and assigns a role to it. If the role assignment changes then the permission level of API keys also change.                                                                                      |
| Radius      | Each connection represents a single env0 organization.                                                                                                                                                                                               |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/env0.spc`)<br />2. Credentials specified in environment variables, e.g., `ENV0_API_KEY`, `ENV0_API_SECRET`.<br />3. Credentials from the env0 CLI configuration file. |

### Configuration

Installing the latest env0 plugin will create a config file (`~/.steampipe/config/env0.spc`) with a single connection named `env0`:

Configure your account details in `~/.steampipe/config/env0.spc`:

```hcl
connection "env0" {
  plugin = "env0"

  # For setting API key and secret see instructions at https://docs.env0.com/reference/authentication
  # `api_key`: The API key for the env0 account. (Required)
  # This can also be set via the `ENV0_API_KEY` environment variable.
  # api_key = "asdpoblfth8acbd"

  # `api_secret`: The API secret of the env0 account. (Required)
  # This can also be set via the `ENV0_API_SECRET` environment variable.
  # api_secret = "LjatOxDqNN9iKH1sLn14TojGkuH3GQAx"

  # If no credentials are specified, the plugin will use env0 CLI configuration, if it has been installed locally.
}
```

Alternatively, you can also use the standard env0 environment variables to obtain credentials **only if other arguments (`api_key` and `api_secret`) are not specified** in the connection:

```sh
export ENV0_API_KEY=asdpoblfth8acbd
export ENV0_API_SECRET=LjatOxDqNN9iKH1sLn14TojGkuH3GQAx
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-env0
- Community: [Slack Channel](https://steampipe.io/community/join)