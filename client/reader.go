package client

import (
	config "sigs.k8s.io/cluster-api/cmd/clusterctl/client/config"
)

func CreateReaderWithConfigs(configs map[string]string) *config.MemoryReader {
	reader := config.NewMemoryReader()
	for key, value := range configs {
		reader.Set(key, value)
	}
	// reader.Set("OPENSTACK_IMAGE_NAME", "image")
	// reader.Set("OPENSTACK_EXTERNAL_NETWORK_ID", "networkid")
	// reader.Set("OPENSTACK_DNS_NAMESERVERS", "8.8.8.8")
	// reader.Set("OPENSTACK_SSH_KEY_NAME", "testssh")
	// reader.Set("OPENSTACK_CLOUD_CACERT_B64", "cacert_b64")
	// reader.Set("OPENSTACK_CLOUD_PROVIDER_CONF_B64", "conf_64")
	// reader.Set("OPENSTACK_CLOUD_YAML_B64", "yaml64")
	// reader.Set("OPENSTACK_FAILURE_DOMAIN", "failure_domain")
	// reader.Set("OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR", "machine")
	// reader.Set("OPENSTACK_NODE_MACHINE_FLAVOR", "node")
	// reader.Set("OPENSTACK_CLOUD", "openstack")
	return reader
}
