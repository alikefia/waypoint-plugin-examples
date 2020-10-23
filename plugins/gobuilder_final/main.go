package main

import (
	"github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final/builder"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	// sdk.Main allows you to register the components which should
	// be included in your plugin
	// Main sets up all the go-plugin requirements
	sdk.Main(sdk.WithComponents(
		&builder.Builder{},
	))
}
