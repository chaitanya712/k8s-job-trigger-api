package utils

import (
	"context"
	"k8s.io/client-go/kubernetes"
)

// FetchContainers extracts container names from a selected CronJob
func FetchContainers(clientset *kubernetes.Clientset, namespace, cronJobName string) ([]string, error) {
	cronJob, err := clientset.BatchV1().CronJobs(namespace).Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var containerNames []string
	for _, container := range cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		containerNames = append(containerNames, container.Name)
	}
	return containerNames, nil
}
