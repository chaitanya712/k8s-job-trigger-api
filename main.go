package main

import (
	"github.com/gin-gonic/gin"
	"k8s-job-trigger-api/routes"
)

func main() {
	r := gin.Default()

	// Initialize API routes
	routes.RegisterRoutes(r)

	r.Run(":8080") // Run on port 8080
}
