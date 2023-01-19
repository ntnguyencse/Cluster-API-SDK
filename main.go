package main

import (
	"fmt"
	"os"

	"github.com/ntnguyencse/cluster-api-sdk/client"
	// clusterclient "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
	kubernetesclient "github.com/ntnguyencse/cluster-api-sdk/kubernetes-client"
)

var kubeconfigFile = "/home/dcn/github/cluster-api-sdk/capi"

func init() {
	fmt.Println("Init function")
	fmt.Println("KUBECONFIG", os.Getenv("KUBECONFIG"))
}

func main() {
	clientset, _ := kubernetesclient.CreateKubernetesClient(&kubeconfigFile)
	kubernetesclient.GetPods(clientset, "default")
	fmt.Println("Main function")
	var configs = map[string]string{
		"OPENSTACK_IMAGE_NAME":                   "OPENSTACK_IMAGE_NAME",
		"OPENSTACK_EXTERNAL_NETWORK_ID":          "OPENSTACK_EXTERNAL_NETWORK_ID",
		"OPENSTACK_DNS_NAMESERVERS":              "OPENSTACK_DNS_NAMESERVERS",
		"OPENSTACK_SSH_KEY_NAME":                 "OPENSTACK_SSH_KEY_NAME",
		"OPENSTACK_CLOUD_CACERT_B64":             "OPENSTACK_CLOUD_CACERT_B64",
		"OPENSTACK_CLOUD_PROVIDER_CONF_B64":      "OPENSTACK_CLOUD_PROVIDER_CONF_B64",
		"OPENSTACK_CLOUD_YAML_B64":               "OPENSTACK_CLOUD_YAML_B64",
		"OPENSTACK_FAILURE_DOMAIN":               "OPENSTACK_FAILURE_DOMAIN",
		"OPENSTACK_CLOUD":                        "OPENSTACK_CLOUD",
		"OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR": "OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR",
		"OPENSTACK_NODE_MACHINE_FLAVOR":          "OPENSTACK_NODE_MACHINE_FLAVOR",
	}
	providerConfigs := client.CreateProviderConfig(client.OPENSTACK, client.OPENSTACK_URL, client.InfrastructureProviderType)
	c, err := client.CreateNewClient(kubeconfigFile, configs, providerConfigs)
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
