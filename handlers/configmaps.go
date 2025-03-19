package handlers

import (
	"k8s-job-trigger-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetConfigMaps fetches ConfigMaps for a selected container
func GetConfigMaps(c *gin.Context) {
	kubeconfig := c.Query("kubeconfig")
	namespace := c.Query("namespace")
	cronJob := c.Query("cronjob")
	container := c.Query("container")

	clientset, err := utils.GetK8sClient(kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load kubeconfig"})
		return
	}

	configMaps, err := utils.FetchConfigMaps(clientset, namespace, cronJob, container)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ConfigMaps"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"configMaps": configMaps})
}
