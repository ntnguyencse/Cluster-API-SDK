package client

import (
	"fmt"

	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	client "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
	cluster "sigs.k8s.io/cluster-api/cmd/clusterctl/client/cluster"
	config "sigs.k8s.io/cluster-api/cmd/clusterctl/client/config"
	repository "sigs.k8s.io/cluster-api/cmd/clusterctl/client/repository"
	cmd "sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
)

var configFile = "/home/ubuntu/cluster-api-sdk/config.yaml"
var kubeconfigFile = "/home/ubuntu/.kube/config"
var _ cmd.Version

var _ client.ClusterClientFactoryInput

type Client struct {
	ConfigClient     config.Client
	ProviderClient   config.Provider
	RepositoryClient repository.Client
	ClusterClient    cluster.Client
	Client           client.Client
	Kubeconfig       cluster.Kubeconfig
}

var Clusterctl Client

func CreateNewClient(path string) (client.Client, error) {
	// Inject config
	var err error
	Clusterctl.ConfigClient, err = config.New(configFile)
	if err != nil {
		fmt.Println("Error create configClient", err)
		return nil, err
	}
	client.InjectConfig(Clusterctl.ConfigClient)
	// Inject Repository Factory
	Clusterctl.ProviderClient = config.NewProvider(
		config.OpenStackProviderName,
		"https://github.com/kubernetes-sigs/cluster-api-provider-openstack/releases/latest/infrastructure-components.yaml",
		clusterctlv1.InfrastructureProviderType,
	)
	Clusterctl.RepositoryClient, err = repository.New(Clusterctl.ProviderClient, Clusterctl.ConfigClient, repository.InjectYamlProcessor(nil))
	if err != nil {
		fmt.Println("Error create repository client", err)
		return nil, err
	}

	Clusterctl.Kubeconfig = cluster.Kubeconfig{Path: path}
	Clusterctl.ClusterClient = cluster.New(Clusterctl.Kubeconfig, Clusterctl.ConfigClient)
	// repoClientFactoryInput := client.RepositoryClientFactoryInput{Provider: provider, Processor: }
	// client.InjectRepositoryFactory(repoClient)
	Clusterctl.Client, err = client.New(configFile,
		client.InjectConfig(Clusterctl.ConfigClient))

	if err != nil {
		fmt.Println("Error create Cluster client", err)
		return nil, err
	}
	return Clusterctl.Client, nil
}
