package client

import (
	"fmt"

	// clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	client "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
	cluster "sigs.k8s.io/cluster-api/cmd/clusterctl/client/cluster"
	config "sigs.k8s.io/cluster-api/cmd/clusterctl/client/config"
	repository "sigs.k8s.io/cluster-api/cmd/clusterctl/client/repository"
	cmd "sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
)

var configFile = "/home/dcn/github/cluster-api-sdk/config.yaml"
var DefaultKubeconfigFile = "/home/dcn/github/cluster-api-sdk/capi"
var _ cmd.Version

var _ client.ClusterClientFactoryInput

// type Kubeconfig cluster.Kubeconfig
type Client struct {
	ConfigClient              config.Client
	ProviderClient            config.Provider
	RepositoryClient          repository.Client
	ClusterClient             cluster.Client
	Client                    client.Client
	Kubeconfig                cluster.Kubeconfig
	ClusterClientFactoryInput client.ClusterClientFactoryInput
}

var Clusterctl Client

// GetKubeconfigOptions carries all the options supported by GetKubeconfig.
type GetKubeconfigOptions client.GetKubeconfigOptions

// type GetKubeconfigOptions struct {
// 	// Kubeconfig defines the kubeconfig to use for accessing the management cluster. If empty,
// 	// default rules for kubeconfig discovery will be used.
// 	Kubeconfig Kubeconfig

// 	// Namespace is the namespace in which secret is placed.
// 	Namespace string

//		// WorkloadClusterName is the name of the workload cluster.
//		WorkloadClusterName string
//	}

// ProviderRepositorySourceOptions defines the options to be used when reading a workload cluster template
// from a provider repository.
type ProviderRepositorySourceOptions client.ProviderRepositorySourceOptions

// type ProviderRepositorySourceOptions struct {
// 	// InfrastructureProvider to read the workload cluster template from. If unspecified, the default
// 	// infrastructure provider will be used if no other sources are specified.
// 	InfrastructureProvider string

// 	// Flavor defines The workload cluster template variant to be used when reading from the infrastructure
// 	// provider repository. If unspecified, the default cluster template will be used.
// 	Flavor string
// }

func CreateNewClient(path string, configs map[string]string, providerConfigs Provider) (Client, error) {
	// Inject config
	// Create Reader

	var err error
	reader := CreateReaderWithConfigs(configs)
	Clusterctl.ConfigClient, err = config.New(configFile, config.InjectReader(reader))
	if err != nil {
		fmt.Println("Error create configClient", err)
		return Clusterctl, err
	}
	client.InjectConfig(Clusterctl.ConfigClient)
	// Inject Repository Factory
	// Config for Provider in here
	Clusterctl.ProviderClient = config.NewProvider(
		providerConfigs.Name,
		providerConfigs.Url,
		providerConfigs.ProviderType,
	)
	//
	Clusterctl.RepositoryClient, err = repository.New(Clusterctl.ProviderClient, Clusterctl.ConfigClient, repository.InjectYamlProcessor(nil))
	if err != nil {
		fmt.Println("Error create repository client", err)
		return Clusterctl, err
	}
	repositoryClientFactory := func(provider config.Provider, configClient config.Client, options ...repository.Option) (repository.Client, error) {
		return Clusterctl.RepositoryClient, nil
	}

	// Convert to client.Config to cluster.Config because of compiler's complaining
	Clusterctl.Kubeconfig = cluster.Kubeconfig{Path: path}
	Clusterctl.ClusterClient = cluster.New(Clusterctl.Kubeconfig, Clusterctl.ConfigClient,
		cluster.InjectRepositoryFactory(repositoryClientFactory),
	)

	// Create clusterClientFactory to override default clusterClientFactory
	clusterClientFactory := func(i client.ClusterClientFactoryInput) (cluster.Client, error) {
		return Clusterctl.ClusterClient, nil
	}
	// // Create Proxy to override default proxy
	// cluster.Proxy.NewClient()
	// Create clusterctl Client
	Clusterctl.Client, err = client.New(configFile,
		client.InjectConfig(Clusterctl.ConfigClient),
		client.InjectClusterClientFactory(clusterClientFactory))

	if err != nil {
		fmt.Println("Error create Cluster client", err)
		return Clusterctl, err
	}
	return Clusterctl, nil
}

func (c *Client) GetKubeconfig(WorkloadClusterName string, Namespace string) (string, error) {
	clientKubeconfig := client.Kubeconfig{Path: c.Kubeconfig.Path}
	options := client.GetKubeconfigOptions{
		Kubeconfig:          clientKubeconfig,
		Namespace:           Namespace,
		WorkloadClusterName: WorkloadClusterName,
	}
	kubeconfig, err := c.Client.GetKubeconfig(options)
	if err != nil {
		fmt.Println("Error when get Kubeconfig ", err)
		return "error", err
	}
	return kubeconfig, nil
}

// GetClusterTemplate Function
// clusterName: name of cluster which will be generated
// ....
// targetNamespace: namespace cluster will be deployed
// infrastructureProvider: infrastructure where cluster will run on
// flavor: postfix for template cluster
func (c *Client) GetClusterTemplate(clusterName string, targetNamespace string, templateUrl string) (string, error) {

	// clientKubeconfig := client.Kubeconfig{Path: c.Kubeconfig.Path, Context: c.Kubeconfig.Context}
	// providerRepositorySourceOptions := client.ProviderRepositorySourceOptions{
	// 	InfrastructureProvider: infrastructureProvider,
	// 	Flavor:                 flavor,
	// }
	urlSource := client.URLSourceOptions{
		URL: templateUrl,
	}
	// Options for generate cluster yaml file
	getClusterTemplateOptions := client.GetClusterTemplateOptions{
		// Kubeconfig: clientKubeconfig,
		// ProviderRepositorySource: &providerRepositorySourceOptions,
		URLSource:       &urlSource,
		TargetNamespace: targetNamespace,
		ClusterName:     clusterName,
	}
	template, err := c.Client.GetClusterTemplate(getClusterTemplateOptions)
	if err != nil {
		fmt.Println("Error when get cluster template", err)
		return "", err
	}

	yamlFile, _ := template.Yaml()
	return string(yamlFile), nil

}
