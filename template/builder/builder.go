package builder

import (
	"context"
	"fmt"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type BuildConfig struct {
	Directory string `hcl:"directory,optional"`
}

type Builder struct {
	config BuildConfig
}

// Config Implements the Waypoint Configurable interface
// Waypoint calls this method before parsing the config inside the use stanza.
//
// It expects a reference to a HCL annotated struct to be returned which will
// be used when de-serialzing the config
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// ConfigSet implements the Waypoint ConfigurableNotify interface.
// Waypoint calls this method after it has deserialized the config to
// the interface returned from the Config method.
func (b *Builder) ConfigSet(config interface{}) error {
	c, ok := config.(*BuildConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *BuildConfig as parameter")
	}

	// validate the config
	if c.Directory == "" {
		return fmt.Errorf("Directory must be set to a valid directory")
	}

	return nil
}

// BuildFunc implements the Builder interface
// Waypoint expects a function to be returned from this method which
// will be called during the build phase of the lifecycle.
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.build
}

// A BuildFunc does not have a strict signature, you define the parameters
// you need based on the available parameters that the Waypoint SDK provides.
// Waypoint automatically injects the parameters specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - *datadir.Project
// - *datadir.App
// - *datadir.Component
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet
//
// The output parameters for BuildFunc must be a Struct which can
// be serialized to Protocol Buffers binary format and an error.
//
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (b *Builder) build(ctx context.Context, ui terminal.UI) (*Binary, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	return &Binary{}, nil
}
