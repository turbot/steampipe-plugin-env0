![image](https://hub.steampipe.io/images/plugins/turbot/env0-social-graphic.png)

# env0 Plugin for Steampipe

Use SQL to query models, completions and more from env0.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/env0)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/env0/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-env0/issues)

## Quick start

### Install

Download and install the latest env0 plugin:

```bash
steampipe plugin install env0
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/env0#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/env0#configuration).

Configure your account details in `~/.steampipe/config/env0.spc`:

```hcl
connection "env0" {
  plugin = "env0"

  # Authentication information
  api_key = "asdpoblfth8acbd" 
  api_secret = "LjatOxDqNN9iKH1sLn14TojGkuH3GQAx"
}
```

Or through environment variables:

```sh
export ENV0_API_KEY=asdpoblfth8acbd
export ENV0_API_SECRET=LjatOxDqNN9iKH1sLn14TojGkuH3GQAx
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-env0.git
cd steampipe-plugin-env0
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/env0.spc
```

Try it!

```
steampipe query
> .inspect env0
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-env0/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-env0/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [env0 Plugin](https://github.com/turbot/steampipe-plugin-env0/labels/help%20wanted)
