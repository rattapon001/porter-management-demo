package main

import (
	"github.com/gin-gonic/gin"
	jobRouter "github.com/rattapon001/porter-management-demo/api/v1/routers/job"
)

func main() {
	router := gin.Default()
	port := "8080"
	jobRouter.InitJobRouter(router)
	router.Run(":" + port)
}
