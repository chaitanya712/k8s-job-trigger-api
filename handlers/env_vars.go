package handlers

import (
	"k8s-job-trigger-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEnvVars fetches environment variables from the ConfigMap
func GetEnvVars(c *gin.Context) {
	kubeconfig := c.Query("kubeconfig")
	namespace := c.Query("namespace")
	cronJob := c.Query("cronjob")
	container := c.Query("container")
	configMap := c.Query("configmap")

	clientset, err := utils.GetK8sClient(kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load kubeconfig"})
		return
	}

	envVars, err := utils.FetchEnvVars(clientset, namespace, cronJob, container, configMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch environment variables"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"envVars": envVars})
}
