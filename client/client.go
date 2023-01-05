package client

import (
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

var (
	ConfigFile     string
	kubeConfigPath string
	kubeConfig     string
)

func New() (client.Client, error) {
	return client.New(configFile)
}

func NewClusterctlClient(clusterctlConfigFile string, kubeCfgPath string) (client.Client, error) {
	clusterctlClient, err := client.New(clusterctlConfigFile)
	if err != nil {
		return nil, err
	}
	// kubeConfigPath = kubeCfgPath
	kubeConfig, err = readKubeConfigFile(kubeCfgPath)
	if err != nil {
		return nil, err
	}
	return clusterctlClient, err
}
