package utils

import (
	"context"
	
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1")

// FetchEnvVars retrieves environment variables from the ConfigMap
func FetchEnvVars(clientset *kubernetes.Clientset, namespace, cronJobName, containerName, configMapName string) ([]map[string]string, error) {
	configMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	envVars := []map[string]string{}
	for key, value := range configMap.Data {
		envVars = append(envVars, map[string]string{"name": key, "value": value})
	}
	return envVars, nil
}
