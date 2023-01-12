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

var configFile = "/home/ubuntu/cluster-api-sdk/config"
var kubeconfigFile = "/home/ubuntu/.kube/config"
var _ cmd.Version

var _ client.ClusterClientFactoryInput

func CreateNewClient(path string) (client.Client, error) {
	// Inject config
	configClient, err := config.New(configFile)
	if err != nil {
		fmt.Println("Error create configClient")
	}
	client.InjectConfig(configClient)
	// Inject Repository Factory
	provider := config.NewProvider(
		config.OpenStackProviderName,
		"https://github.com/kubernetes-sigs/cluster-api-provider-openstack/releases/latest/infrastructure-components.yaml",
		clusterctlv1.InfrastructureProviderType,
	)
	repositoryClient, err1 := repository.New(provider, configClient, repository.InjectYamlProcessor(nil))
	if err1 != nil {
		fmt.Println("Error create repository client")
		_ = repositoryClient
	}

	kubeConfig := cluster.Kubeconfig{Path: kubeconfigFile}
	clusterClient := cluster.New(kubeConfig, configClient)
	// repoClientFactoryInput := client.RepositoryClientFactoryInput{Provider: provider, Processor: }
	// client.InjectRepositoryFactory(repoClient)
	client, err2 := client.New(configFile,
		client.InjectConfig(configClient))
}
