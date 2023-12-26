package main

import (
	"github.com/gin-gonic/gin"
	porter_router "github.com/rattapon001/porter-management-demo/api/v1/routers/porter"
)

func main() {
	router := gin.Default()
	port := "8081"
	porter_router.InitPorterRouter(router)
	router.Run(":" + port)
}
