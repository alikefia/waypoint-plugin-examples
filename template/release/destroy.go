package release

import (
	"context"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// DestroyFunc implements the Waypoint Destroyer interface.
// Waypoint expects a function to be returned from this method
// which is called during the destroy phase of the lifecycle.
func (rm *ReleaseManager) DestroyFunc() interface{} {
	return rm.destroy
}

// A DestroyFunc does not have a strict signature, you define the parameters
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
// The output parameters for DestroyFunc must be a Struct which can
// be serialized to Protocol Buffers binary format and an error.
//
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (rm *ReleaseManager) destroy(ctx context.Context, ui terminal.UI, release *Release) error {
	return nil
}
