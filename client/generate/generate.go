package generate

import (
	_ "fmt"
	_ "os"

	_ "github.com/pkg/errors"

	// "github.com/spf13/cobra"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
	// "sigs.k8s.io/cluster-api/cmd"
	wrapper "github.com/ntnguyencse/cluster-api-sdk/client"
)

type GenerateClusterOptions struct {
	kubeconfig             string
	kubeconfigContext      string
	flavor                 string
	infrastructureProvider string

	targetNamespace          string
	kubernetesVersion        string
	controlPlaneMachineCount int64
	workerMachineCount       int64

	url                string
	configMapNamespace string
	configMapName      string
	configMapDataKey   string

	listVariables bool
}

func init() {

}

// # Generates a yaml file for creating workload clusters using
// # the pre-installed infrastructure and bootstrap providers.
// clusterctl generate cluster my-cluster
// # Generates a yaml file for creating workload clusters using
// # a specific version of the AWS infrastructure provider.
// clusterctl generate cluster my-cluster --infrastructure=aws:v0.4.1
// # Generates a yaml file for creating workload clusters in a custom namespace.
// clusterctl generate cluster my-cluster --target-namespace=foo
// # Generates a yaml file for creating workload clusters with a specific Kubernetes version.
// clusterctl generate cluster my-cluster --kubernetes-version=v1.19.1
// # Generates a yaml file for creating workload clusters with a
// # custom number of nodes (if supported by the provider's templates).
// clusterctl generate cluster my-cluster --control-plane-machine-count=3 --worker-machine-count=10
// # Generates a yaml file for creating workload clusters using a template stored in a ConfigMap.
// clusterctl generate cluster my-cluster --from-config-map MyTemplates
// # Generates a yaml file for creating workload clusters using a template from a specific URL.
// clusterctl generate cluster my-cluster --from https://github.com/foo-org/foo-repository/blob/main/cluster-template.yaml
// # Generates a yaml file for creating workload clusters using a template stored locally.
// clusterctl generate cluster my-cluster --from ~/workspace/cluster-template.yaml
// # Prints the list of variables required by the yaml file for creating workload cluster.
// clusterctl generate cluster my-cluster --list-variables`),

func GenerateKubernetesCluster(clusterName string, options GenerateClusterOptions) error {
	var gc = &GenerateClusterOptions{}
	c, err := client.New(wrapper.ConfigFile)

	c.GetProvidersConfig()
	if err != nil {
		return err
	}

	templateOptions := client.GetClusterTemplateOptions{
		Kubeconfig:        client.Kubeconfig{Path: gc.kubeconfig, Context: gc.kubeconfigContext},
		ClusterName:       clusterName,
		TargetNamespace:   gc.targetNamespace,
		KubernetesVersion: gc.kubernetesVersion,
		ListVariablesOnly: gc.listVariables,
	}
	// Assign value to command's options.
	templateOptions.ControlPlaneMachineCount = &options.controlPlaneMachineCount
	templateOptions.WorkerMachineCount = &options.workerMachineCount
	// templateOptions.Flavor = options.flavor
	templateOptions.TargetNamespace = options.targetNamespace
	// templateOptions.InfrastructureProvider = options.infrastructureProvider
	templateOptions.KubernetesVersion = options.kubernetesVersion
	return nil
}
