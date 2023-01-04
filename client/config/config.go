package config

import (
	"os"

	_ "github.com/pkg/errors"
	_ "sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
)

func NewConfig(path string) {

	return
}

func GetKubeConfig(path string) ([]byte, error) {
	kubeconfigFile, err := os.ReadFile("/tmp/dat")
	if err != nil {
		return []byte("Error"), err
	}
	return kubeconfigFile, nil
}
