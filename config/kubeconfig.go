package config

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// LoadKubeconfig parses kubeconfig from the provided path or data
func LoadKubeconfig(kubeconfigData string) (*api.Config, error) {
	config, err := clientcmd.Load([]byte(kubeconfigData))
	if err != nil {
		return nil, err
	}
	return config, nil
}
