package kubernetesclient

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	k8scmd "k8s.io/kubectl/pkg/cmd"
	k8scmdplugin "k8s.io/kubectl/pkg/cmd/plugin"

	// "github.com/ntnguyencse/cluster-api-sdk/kubernetes-client/genericiooptions"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	// "k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

type ClientSet *kubernetes.Clientset

// var clientSet *kubernetes.Clientset

func CreateKubernetesClient(kubeConfigPath *string) (ClientSet, error) {
	var kubeconfig *string

	kubeconfig = flag.String("kubeconfig", *kubeConfigPath, "absolute path to the kubeconfig file")

	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Build config format failed")
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to create kubernetes client")
		return nil, err
	}
	return clientset, err
}
func GetPods(clientset *kubernetes.Clientset, namespace string) {

	// crds, err := clientset.RESTClient().Get().Namespace("default").
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		pod := "example-xxxxx"
		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}
func KubectlApplyYamlFile(yamlFilePath string) {
	arg := []string{"kubectl", "apply"}
	var defaultConfigFlags = genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
	kubectlOptions := k8scmd.KubectlOptions{
		PluginHandler: k8scmd.NewDefaultPluginHandler(k8scmdplugin.ValidPluginFilenamePrefixes),
		Arguments:     arg,
		ConfigFlags:   defaultConfigFlags,
		IOStreams:     genericclioptions.IOStreams{In: strings.NewReader(""), Out: os.Stdout, ErrOut: os.Stderr},
	}
	// k8scmd.NewDefaultKubectlCommandWithArgs(kubectlOptions)
	cmd := k8scmd.NewKubectlCommand(kubectlOptions)
	_ = cmd
	// fmt.Println("Print cmd: ", cmd)
	kubeConfigFlags := kubectlOptions.ConfigFlags
	if kubeConfigFlags == nil {
		kubeConfigFlags = defaultConfigFlags
	}
	// kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	// matchVersionKubeConfigFlags.AddFlags(flags)
	// Updates hooks to add kubectl command headers: SIG CLI KEP 859.
	// addCmdHeaderHooks(cmds, kubeConfigFlags)
	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)

	applyIOStream, _, outbuff, _ := genericclioptions.NewTestIOStreams()
	applyCmd := apply.NewCmdApply("kubectl", f, applyIOStream)

	applyCmd.Flags().Set("filename", yamlFilePath)
	// Enable Debug Flags in cobra commands
	// applyCmd.DebugFlags()
	applyCmd.Run(applyCmd, []string{})

	fmt.Println("End of function", "out", outbuff)
}
func KubectlApplyDefault(yamlString *string) {
	arg := []string{"kubectl", "apply"}
	var defaultConfigFlags = genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
	kubectlOptions := k8scmd.KubectlOptions{
		PluginHandler: k8scmd.NewDefaultPluginHandler(k8scmdplugin.ValidPluginFilenamePrefixes),
		Arguments:     arg,
		ConfigFlags:   defaultConfigFlags,
		IOStreams:     genericclioptions.IOStreams{In: strings.NewReader(*yamlString), Out: os.Stdout, ErrOut: os.Stderr},
	}
	// k8scmd.NewDefaultKubectlCommandWithArgs(kubectlOptions)
	cmd := k8scmd.NewKubectlCommand(kubectlOptions)
	// fmt.Println("Print cmd: ", cmd)
	kubeConfigFlags := kubectlOptions.ConfigFlags
	if kubeConfigFlags == nil {
		kubeConfigFlags = defaultConfigFlags
	}
	// kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	// matchVersionKubeConfigFlags.AddFlags(flags)
	// Updates hooks to add kubectl command headers: SIG CLI KEP 859.
	// addCmdHeaderHooks(cmds, kubeConfigFlags)
	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	// applyIOStream := genericclioptions.IOStreams{
	// 	In:     strings.NewReader(*yamlString),
	// 	Out:    &bytes.Buffer{},
	// 	ErrOut: &bytes.Buffer{},
	// }
	applyIOStream, _, outbuff, _ := genericclioptions.NewTestIOStreams()
	applyCmd := apply.NewCmdApply("kubectl", f, applyIOStream)

	// applyCmd.SetArgs([]string{"-f", "/home/ubuntu/l-kaas/Cluster-API-SDK/test.yaml"})

	applyCmd.Flags().Set("filename", "/home/ubuntu/l-kaas/Cluster-API-SDK/test.yaml")
	// applyCmd.Flags().Set("dry-run", "client")
	// applyCmd.Flags().Set("server-side", "true")
	// applyCmd.Flags().Set("output", "json")
	// applyCmd.Flags().Set("prune", "true")

	// Create a local builder...
	// builder := resource.NewLocalBuilder().
	// 	// Configure with a scheme to get typed objects in the versions registered with the scheme.
	// 	// As an alternative, could call Unstructured() to get unstructured objects.
	// 	WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
	// 	// Provide input via a Reader.
	// 	// As an alternative, could call Path(false, "/path/to/file") to read from a file.
	// 	Stream(bytes.NewBufferString(*yamlString), "input").
	// 	// Flatten items contained in List objects
	// 	Flatten().
	// 	// Accumulate as many items as possible
	// 	ContinueOnError()

	// // Run the builder
	// result := builder.Do()

	// applyCmd.DebugFlags()
	applyCmd.Run(applyCmd, []string{})

	// err := applyCmd.Execute()
	_ = cmd
	// cmd.AddCommand(applyCmd)

	// err := cmd.Execute()

	fmt.Println("End of function", "out", outbuff)

}
func KubectlApply(yamlString *string) {
	kubeConfig := clientcmdapi.NewConfig()

	// tf := cmdtesting.NewTestFactory().WithNamespace("test")

	ioStreams := genericclioptions.IOStreams{
		In:     strings.NewReader(*yamlString),
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	_ = clientcmd.NewDefaultClientConfig(*kubeConfig, nil)
	configFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
	// configFlags.WithCon

	tf := cmdutil.NewFactory(configFlags)
	// tf.WithClientConfig(clientConfig)

	// kubeConfig.Contexts["default"] = &clientcmdapi.Context{Namespace: "bar"}
	cmd := apply.NewCmdApply("kubectl", tf, ioStreams)
	// cmd.Flags().Set("output", "name")
	_ = genericclioptions.NewRecordFlags()
	cmd.Run(cmd, []string{})

}
func ApplyResourceString(kubePath string, yamlString string) error {
	// Create a new Kubernetes clientset
	// clientset, err := CreateKubernetesClient(&kubePath)
	// if err != nil {
	// 	return err
	// }

	// // Create a new "kubectl apply" command
	// ioStreams := genericclioptions.IOStreams{
	// 	In:     bytes.NewBufferString(yamlString),
	// 	Out:    os.Stdout,
	// 	ErrOut: os.Stderr,
	// }
	// resourceConfig := &resource.Info{
	// 	Namespace: "default",
	// 	Source:    &unstructured.Unstructured{Object: map[string]interface{}{}},
	// }
	// if err := resourceConfig.Source.Unmarshal([]byte(yamlString)); err != nil {
	// 	return err
	// }
	// // Create a new apply command
	// cmd := apply.NewCmdApply("kubectl", ioStreams)
	// cmd.SetArgs([]string{"-f", "-"})

	// // Run the apply command

	// if err != nil {
	// 	return fmt.Errorf("failed to apply resource: %v", err)
	// }

	return nil
}
func NewRestConfig(kubeconfig string) (*rest.Config, error) {
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return config, nil
}
