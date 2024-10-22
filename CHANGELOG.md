## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#34](https://github.com/turbot/steampipe-plugin-env0/pull/34))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#34](https://github.com/turbot/steampipe-plugin-env0/pull/34))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#29](https://github.com/turbot/steampipe-plugin-env0/pull/29))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#29](https://github.com/turbot/steampipe-plugin-env0/pull/29))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-env0/blob/main/docs/LICENSE). ([#29](https://github.com/turbot/steampipe-plugin-env0/pull/29))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#28](https://github.com/turbot/steampipe-plugin-env0/pull/28))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#12](https://github.com/turbot/steampipe-plugin-env0/pull/12))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#8](https://github.com/turbot/steampipe-plugin-env0/pull/8))
- Recompiled plugin with Go version `1.21`. ([#8](https://github.com/turbot/steampipe-plugin-env0/pull/8))

## v0.0.1 [2023-05-19]

_What's new?_

- New tables added
  - [env0_api_key](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_api_key)
  - [env0_environment](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_environment)
  - [env0_notification](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_notification)
  - [env0_organization](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_organization)
  - [env0_project](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_project)
  - [env0_role](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_role)
  - [env0_team](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_team)
  - [env0_template](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_template)
  - [env0_user](https://hub.steampipe.io/plugins/turbot/env0/tables/env0_user)
