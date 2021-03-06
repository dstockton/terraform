package main

import (
	"github.com/hashicorp/terraform/builtin/providers/rightscale"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rightscale.Provider,
	})
}
