// Package files contains a component for validating local files.
package main

import (
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/platform"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/registry"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/release"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	sdk.Main(sdk.WithComponents(
		&registry.Registry{},
		&platform.Deploy{},
		&release.Releaser{},
	))
}
