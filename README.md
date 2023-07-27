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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-env0/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [env0 Plugin](https://github.com/turbot/steampipe-plugin-env0/labels/help%20wanted)
