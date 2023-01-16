package main

import (
	"fmt"
	"os"

	"github.com/ntnguyencse/cluster-api-sdk/client"
	// clusterclient "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

var kubeconfigFile = "/home/ubuntu/.kube/config"

func init() {
	fmt.Println("Init function")
	fmt.Println("KUBECONFIG", os.Getenv("KUBECONFIG"))
}

func main() {
	fmt.Println("Main function")
	c, err := client.CreateNewClient(kubeconfigFile)
	fmt.Println("Created client")
	if err != nil {
		fmt.Println("Error when create client", err)
	}
	provider, err := c.Client.GetProvidersConfig()
	fmt.Println("Get Provider Informations")
	fmt.Println(provider[0])
	fmt.Println("Get KubeConfig")

	kubeCluster, err := c.GetKubeconfig("test1", "default")
	fmt.Println(kubeCluster)
	c.GetClusterTemplate("a", "1.24.8", 3, 3, "test", c.ProviderClient.Name(), "medium")
}
