package main

import (
	"github.com/turbot/steampipe-plugin-env0/env0"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: env0.Plugin})
}
