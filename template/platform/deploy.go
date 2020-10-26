package platform

import (
	"context"

	"github.com/hashicorp/waypoint-plugin-examples/template/registry"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// Config is used by Waypoint when serializing the config defined
// in the "use" stanza
//
//use "myplugin" {
//	region = "my name"
//}
type Config struct {
	Region string `hcl:"region,optional"`
}

// Platform defines a Waypoint component which can be used
// during the deploy phase.
type Platform struct {
	config Config
}

// Config Implements the Waypoint Configurable interface
// Waypoint calls this method before parsing the config inside the use stanza.
//
// It expects a reference to a HCL annotated struct to be returned which will
// be used when de-serialzing the config
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// ConfigSet implements the Waypoint ConfigurableNotify interface.
// Waypoint calls this method after it has deserialized the config to
// the interface returned from the Config method.
func (p *Platform) ConfigSet(config interface{}) error {
	//c, ok := config.(*Config)
	//if !ok {
	//	// The Waypoint SDK should ensure this never gets hit
	//	return fmt.Errorf("Expected *DeployConfig as parameter")
	//}

	return nil
}

// DeployFunc implements the Platform interface
// Waypoint expects a function to be returned from this method which
// will be called during the deploy phase of the lifecycle.
func (p *Platform) DeployFunc() interface{} {
	// return a function which will be called by Waypoint
	return p.deploy
}

// A DeployFunc does not have a strict signature, you define the parameters
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
// The output parameters for DeployFunc must be a Struct which can
// be serialized to Protocol Buffers binary format and an error.
//
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (p *Platform) deploy(ctx context.Context, ui terminal.UI, artifact *registry.Artifact) (*Deployment, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Deploy application")

	return &Deployment{}, nil
}
