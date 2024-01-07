package job_router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitJobRouter(router *gin.Engine, db *mongo.Client) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello job",
		})
	})

	jobRouter := router.Group("/jobs")
	{
		jobRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello job",
			})
		})
	}
}
