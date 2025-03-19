package utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// FetchCronJobs retrieves all CronJobs in a given namespace
func FetchCronJobs(clientset *kubernetes.Clientset, namespace string) ([]map[string]interface{}, error) {
	cronjobs, err := clientset.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var cronJobList []map[string]interface{}
	for _, cj := range cronjobs.Items {
		containers := []string{}
		for _, container := range cj.Spec.JobTemplate.Spec.Template.Spec.Containers {
			containers = append(containers, container.Name)
		}

		cronJobList = append(cronJobList, map[string]interface{}{
			"name":       cj.Name,
			"containers": containers,
		})
	}
	return cronJobList, nil
}
