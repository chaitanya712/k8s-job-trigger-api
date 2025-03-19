package handlers

import (
	"k8s-job-trigger-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetContainers fetches containers from a selected CronJob
func GetContainers(c *gin.Context) {
	kubeconfig := c.Query("kubeconfig")
	namespace := c.Query("namespace")
	cronJob := c.Query("cronjob")

	clientset, err := utils.GetK8sClient(kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load kubeconfig"})
		return
	}

	containers, err := utils.FetchContainers(clientset, namespace, cronJob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch containers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"containers": containers})
}
