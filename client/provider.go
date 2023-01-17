package client

import (
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
)

const (
	// CoreProviderType is a type reserved for Cluster API core repository.
	CoreProviderType = clusterctlv1.ProviderType("CoreProvider")

	// BootstrapProviderType is the type associated with codebases that provide
	// bootstrapping capabilities.
	BootstrapProviderType = clusterctlv1.ProviderType("BootstrapProvider")

	// InfrastructureProviderType is the type associated with codebases that provide
	// infrastructure capabilities.
	InfrastructureProviderType = clusterctlv1.ProviderType("InfrastructureProvider")

	// ControlPlaneProviderType is the type associated with codebases that provide
	// control-plane capabilities.
	ControlPlaneProviderType = clusterctlv1.ProviderType("ControlPlaneProvider")

	// IPAMProviderType is the type associated with codebases that provide
	// IPAM capabilities.
	IPAMProviderType = clusterctlv1.ProviderType("IPAMProvider")

	// RuntimeExtensionProviderType is the type associated with codebases that provide
	// runtime extensions.
	RuntimeExtensionProviderType = clusterctlv1.ProviderType("RuntimeExtensionProvider")

	// ProviderTypeUnknown is used when the type is unknown.
	ProviderTypeUnknown = clusterctlv1.ProviderType("")
)
const (
	OPENSTACK = "openstack"
)
const (
	OPENSTACK_URL = "https://github.com/kubernetes-sigs/cluster-api-provider-openstack/releases/download/v0.6.4/infrastructure-components.yaml"
)

type Provider struct {
	Name         string
	Url          string
	ProviderType clusterctlv1.ProviderType
}

func CreateProviderConfig(name string, url string, providerType clusterctlv1.ProviderType) Provider {
	return Provider{
		Name:         name,
		Url:          url,
		ProviderType: providerType,
	}
}
