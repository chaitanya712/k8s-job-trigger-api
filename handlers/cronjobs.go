package handlers

import (
	"k8s-job-trigger-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCronJobs fetches CronJobs for a namespace
func GetCronJobs(c *gin.Context) {
	kubeconfig := c.Query("kubeconfig")
	namespace := c.Query("namespace")

	clientset, err := utils.GetK8sClient(kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load kubeconfig"})
		return
	}

	cronJobs, err := utils.FetchCronJobs(clientset, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch CronJobs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cronjobs": cronJobs})
}
