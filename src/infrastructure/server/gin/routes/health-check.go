package routes

import "github.com/gin-gonic/gin"

func healthCheckResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HealthCheck(r *gin.Engine) {
	account := r.Group("/health-check")
	{
		account.GET("", healthCheckResponse)
	}
}
