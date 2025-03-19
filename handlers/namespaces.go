package handlers

import (
	"k8s-job-trigger-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNamespaces fetches all namespaces
func GetNamespaces(c *gin.Context) {
	kubeconfig := c.Query("kubeconfig")
	clientset, err := utils.GetK8sClient(kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load kubeconfig"})
		return
	}

	namespaces, err := utils.FetchNamespaces(clientset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch namespaces"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"namespaces": namespaces})
}
