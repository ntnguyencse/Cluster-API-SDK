package main

import (
	"fmt"
	"os"

	"github.com/ntnguyencse/cluster-api-sdk/client"
)

var kubeconfigFile = "/home/ubuntu/.kube/config"

func init() {
	fmt.Println("Init function")
	fmt.Println("KUBECONFIG", os.Getenv("KUBECONFIG"))
}

func main() {
	fmt.Println("Main function")
	Client, err := client.CreateNewClient(kubeconfigFile)
	fmt.Println("Created client")
	if err != nil {
		fmt.Println("Error when create client", err)
	}
	fmt.Println(Client.GetProvidersConfig())
}
