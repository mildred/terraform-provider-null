package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mildred/terraform-provider-null/null"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: null.Provider})
}
