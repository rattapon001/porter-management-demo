package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	jobRouter "github.com/rattapon001/porter-management-demo/api/v1/routers/job"
	job_mongo "github.com/rattapon001/porter-management-demo/internal/job/infra/mongo"
)

func main() {

	err := godotenv.Load("./configs/local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	db := job_mongo.MongoDbInit()
	router := gin.Default()
	port := "8080"
	jobRouter.InitJobRouter(router, db)
	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	router.Run(":" + port)
}
