package main

import (
	"fmt"
	"os"

	"github.com/ntnguyencse/cluster-api-sdk/client"
	// clusterclient "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

var kubeconfigFile = "/home/dcn/github/cluster-api-sdk/capi"

func init() {
	fmt.Println("Init function")
	fmt.Println("KUBECONFIG", os.Getenv("KUBECONFIG"))
}

func main() {
	fmt.Println("Main function")
	var configs = map[string]string{
		"OPENSTACK_IMAGE_NAME":                   "k8s-v1.24.8",
		"OPENSTACK_EXTERNAL_NETWORK_ID":          "network.id",
		"OPENSTACK_DNS_NAMESERVERS":              "8.8.8.8",
		"OPENSTACK_SSH_KEY_NAME":                 "abc",
		"OPENSTACK_CLOUD_CACERT_B64":             "tesst",
		"OPENSTACK_CLOUD_PROVIDER_CONF_B64":      "conf-b64",
		"OPENSTACK_CLOUD_YAML_B64":               "yaml64",
		"OPENSTACK_FAILURE_DOMAIN":               "failure.domain",
		"OPENSTACK_CLOUD":                        "cloud",
		"OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR": "machine",
		"OPENSTACK_NODE_MACHINE_FLAVOR":          "node",
	}
	c, err := client.CreateNewClient(kubeconfigFile, configs)
	fmt.Println("Created client")
	if err != nil {
		fmt.Println("Error when create client", err)
	}
	provider, err := c.Client.GetProvidersConfig()
	fmt.Println("Get Provider Informations")
	fmt.Println(provider[0])
	fmt.Println("Get KubeConfig")

	kubeCluster, err := c.GetKubeconfig("my-cluster", "default")
	fmt.Println(kubeCluster)
	// Infrastructure must include version inside input: Ex openstack v0.6.4
	c.GetClusterTemplate("a", "1.24.8", 3, 3, "test", "openstack:v0.6.4", "")
}
