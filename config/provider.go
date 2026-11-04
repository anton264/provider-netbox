package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "netbox"
	modulePath     = "github.com/anton264/provider-netbox"

	rootGroup           = "netbox.crossplane.io"
	rootGroupNamespaced = "netbox.m.crossplane.io"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// defaultResourceOptions are applied to every NetBox resource, in order.
func defaultResourceOptions() ujconfig.ProviderOption {
	return ujconfig.WithDefaultResourceOptions(
		ExternalNameConfigurations(),
		GroupKindOverrides(),
		References(),
	)
}

// GetProvider returns the cluster-scoped (Crossplane v1 style) provider configuration.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		defaultResourceOptions(),
	)

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced (Crossplane v2 style) provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroupNamespaced),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		defaultResourceOptions(),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}),
	)

	pc.ConfigureResources()
	return pc
}
