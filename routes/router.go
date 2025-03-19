package routes

import (
	"k8s-job-trigger-backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines API routes
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/namespaces", handlers.GetNamespaces)
		api.GET("/cronjobs", handlers.GetCronJobs)
		api.GET("/containers", handlers.GetContainers)
		api.GET("/configmaps", handlers.GetConfigMaps)
		api.GET("/envvars", handlers.GetEnvVars)
	}
}
