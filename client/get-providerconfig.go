package client

import (
	client "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

// type Provider struct {
// 	metav1.TypeMeta   `json:",inline"`
// 	metav1.ObjectMeta `json:"metadata,omitempty"`

// 	// ProviderName indicates the name of the provider.
// 	// +optional
// 	ProviderName string `json:"providerName,omitempty"`

// 	// Type indicates the type of the provider.
// 	// See ProviderType for a list of supported values
// 	// +optional
// 	Type string `json:"type,omitempty"`

// 	// Version indicates the component version.
// 	// +optional
// 	Version string `json:"version,omitempty"`

// 	// WatchedNamespace indicates the namespace where the provider controller is watching.
// 	// If empty the provider controller is watching for objects in all namespaces.
// 	//
// 	// Deprecated: in clusterctl v1alpha4 all the providers watch all the namespaces; this field will be removed in a future version of this API
// 	// +optional
// 	WatchedNamespace string `json:"watchedNamespace,omitempty"`
// }

func GetProvidersConfig() ([]client.Provider, error) {
	clusterctlCLient, err := NewClusterctlClient(configFile, kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return clusterctlCLient.GetProvidersConfig()
}
