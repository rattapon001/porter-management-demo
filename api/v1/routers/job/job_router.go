package job_router

import (
	"github.com/gin-gonic/gin"
)

func InitJobRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello job",
		})
	})
}
