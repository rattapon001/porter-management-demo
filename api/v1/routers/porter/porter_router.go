package porter_router

import (
	"github.com/gin-gonic/gin"
)

func InitPorterRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello porter",
		})
	})
}
