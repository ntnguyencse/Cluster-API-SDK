package client

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/utils/pointer"

	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

// printYamlOutput prints the yaml content of a generated template to stdout.
func printYamlOutput(printer client.YamlPrinter) error {
	yaml, err := printer.Yaml()
	if err != nil {
		return err
	}
	yaml = append(yaml, '\n')

	if _, err := os.Stdout.Write(yaml); err != nil {
		return errors.Wrap(err, "failed to write yaml to Stdout")
	}
	return nil
}
func YamlToString(printer client.YamlPrinter) (string, error) {
	yaml, err := printer.Yaml()
	if err != nil {
		return "error", err
	}
	yaml = append(yaml, '\n')

	return string(yaml), nil
}

// printVariablesOutput prints the expected variables in the template to stdout.
func printVariablesOutput(template client.Template, options client.GetClusterTemplateOptions) error {
	// Decorate the variable map for printing
	variableMap := template.VariableMap()
	var requiredVariables []string
	var optionalVariables []string
	for name := range variableMap {
		if variableMap[name] != nil {
			v := *variableMap[name]
			// Add quotes around any unquoted strings
			if len(v) > 0 && !strings.HasPrefix(v, "\"") {
				v = fmt.Sprintf("%q", v)
				variableMap[name] = &v
			}
		}

		// Fix up default for well-know variables that have a special logic implemented in clusterctl.
		// NOTE: this logic mimics the defaulting rules implemented in client.GetClusterTemplate;
		switch name {
		case "CLUSTER_NAME":
			// Cluster name from the cmd arguments is used instead of template default.
			variableMap[name] = pointer.String(options.ClusterName)
		case "NAMESPACE":
			// Namespace name from the cmd flags or from the kubeconfig is used instead of template default.
			if options.TargetNamespace != "" {
				variableMap[name] = pointer.String(options.TargetNamespace)
			} else {
				variableMap[name] = pointer.String("current Namespace in the KubeConfig file")
			}
		case "CONTROL_PLANE_MACHINE_COUNT":
			// Control plane machine count uses the cmd flag, env variable or a constant is used instead of template default.
			if options.ControlPlaneMachineCount == nil {
				if val, ok := os.LookupEnv("CONTROL_PLANE_MACHINE_COUNT"); ok {
					variableMap[name] = pointer.String(val)
				} else {
					variableMap[name] = pointer.String("1")
				}
			} else {
				variableMap[name] = pointer.String(strconv.FormatInt(*options.ControlPlaneMachineCount, 10))
			}
		case "WORKER_MACHINE_COUNT":
			// Worker machine count uses the cmd flag, env variable or a constant is used instead of template default.
			if options.WorkerMachineCount == nil {
				if val, ok := os.LookupEnv("WORKER_MACHINE_COUNT"); ok {
					variableMap[name] = pointer.String(val)
				} else {
					variableMap[name] = pointer.String("0")
				}
			} else {
				variableMap[name] = pointer.String(strconv.FormatInt(*options.WorkerMachineCount, 10))
			}
		case "KUBERNETES_VERSION":
			// Kubernetes version uses the cmd flag, env variable, or the template default.
			if options.KubernetesVersion != "" {
				variableMap[name] = pointer.String(options.KubernetesVersion)
			} else if val, ok := os.LookupEnv("KUBERNETES_VERSION"); ok {
				variableMap[name] = pointer.String(val)
			}
		}

		if variableMap[name] != nil {
			optionalVariables = append(optionalVariables, name)
		} else {
			requiredVariables = append(requiredVariables, name)
		}
	}
	sort.Strings(requiredVariables)
	sort.Strings(optionalVariables)

	if len(requiredVariables) > 0 {
		fmt.Println("Required Variables:")
		for _, v := range requiredVariables {
			fmt.Printf("  - %s\n", v)
		}
	}

	if len(optionalVariables) > 0 {
		fmt.Println("\nOptional Variables:")
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.FilterHTML)
		for _, v := range optionalVariables {
			fmt.Fprintf(w, "  - %s\t(defaults to %s)\n", v, *variableMap[v])
		}
		if err := w.Flush(); err != nil {
			return err
		}
	}

	fmt.Println()
	return nil
}

// printComponentsAsText prints information about the components to stdout.
func printComponentsAsText(c client.Components) error {
	dir, file := filepath.Split(c.URL())
	// Remove the version suffix from the URL since we already display it
	// separately.
	baseURL, _ := filepath.Split(strings.TrimSuffix(dir, "/"))
	fmt.Printf("Name:               %s\n", c.Name())
	fmt.Printf("Type:               %s\n", c.Type())
	fmt.Printf("URL:                %s\n", baseURL)
	fmt.Printf("Version:            %s\n", c.Version())
	fmt.Printf("File:               %s\n", file)
	fmt.Printf("TargetNamespace:    %s\n", c.TargetNamespace())
	if len(c.Variables()) > 0 {
		fmt.Println("Variables:")
		for _, v := range c.Variables() {
			fmt.Printf("  - %s\n", v)
		}
	}
	if len(c.Images()) > 0 {
		fmt.Println("Images:")
		for _, v := range c.Images() {
			fmt.Printf("  - %s\n", v)
		}
	}
	fmt.Println()

	return nil
}

// visitCommands visits the commands.
func visitCommands(cmd *cobra.Command, fn func(*cobra.Command)) {
	fn(cmd)

	for _, c := range cmd.Commands() {
		visitCommands(c, fn)
	}
}

func readKubeConfigFile(kubeconfigPath string) (string, error) {
	kubeConfigData, err := os.ReadFile(kubeconfigPath)
	if err != nil {
		return "Cant read kubeconfig", err
	}
	return string(kubeConfigData), err
}
