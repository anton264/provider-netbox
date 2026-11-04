package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider. NetBox assigns a numeric ID to every object, and the Terraform
// provider uses that ID for import, so every resource uses
// IdentifierFromProvider: the external name is the NetBox object ID. To adopt
// an existing NetBox object, set the crossplane.io/external-name annotation to
// its ID.
var ExternalNameConfigs = func() map[string]config.ExternalName {
	m := make(map[string]config.ExternalName, len(groupKinds))
	for name := range groupKinds {
		m[name] = config.IdentifierFromProvider
	}
	return m
}()

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually. Only these resources are generated.
func ExternalNameConfigured() []string {
	l := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l = append(l, name+"$")
	}
	return l
}
