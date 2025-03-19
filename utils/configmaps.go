package utils

import (
	"context"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FetchConfigMaps lists ConfigMaps associated with a container
func FetchConfigMaps(clientset *kubernetes.Clientset, namespace, cronJobName, containerName string) ([]string, error) {
	cronJob, err := clientset.BatchV1().CronJobs(namespace).Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var configMapNames []string
	for _, container := range cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			for _, envFrom := range container.EnvFrom {
				if envFrom.ConfigMapRef != nil {
					configMapNames = append(configMapNames, envFrom.ConfigMapRef.Name)
				}
			}
		}
	}
	return configMapNames, nil
}
