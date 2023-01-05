package client

import (
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

type GetKubeconfigOptions struct {
	workloadClusterName string
	kubeconfig          string
	kubeconfigContext   string
	namespace           string
}

func runGetKubeconfig(gk GetKubeconfigOptions) (string, error) {
	c, err := client.New(configFile)
	if err != nil {
		return "Cant create clusterctl client", err
	}

	options := client.GetKubeconfigOptions{
		Kubeconfig:          client.Kubeconfig{Path: gk.kubeconfig, Context: gk.kubeconfigContext},
		WorkloadClusterName: gk.workloadClusterName,
		Namespace:           gk.namespace,
	}

	out, err := c.GetKubeconfig(options)
	if err != nil {
		return "Cant get kubeconfig", err
	}
	// fmt.Println(out)
	return out, err
}
