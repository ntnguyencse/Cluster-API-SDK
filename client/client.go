package client

import (
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

var (
	ConfigFile string
)

func NewClusterctlClient(clusterctlConfigFile string) (client.Client, error) {
	clusterctlClient, err := client.New(clusterctlConfigFile)
	if err != nil {
		return nil, err
	}
	return clusterctlClient, err
}
