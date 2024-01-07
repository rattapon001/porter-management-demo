package job_router

import (
	"github.com/gin-gonic/gin"
	job_handler "github.com/rattapon001/porter-management-demo/api/v1/handlers/job"
	"github.com/rattapon001/porter-management-demo/internal/job/app"
	mongoInfra "github.com/rattapon001/porter-management-demo/internal/job/infra/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitJobRouter(router *gin.Engine, db *mongo.Client) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello job",
		})
	})
	coll := db.Database("porter").Collection("jobs")
	jobRepository := mongoInfra.NewJobMongoRepository(coll)
	JobService := app.NewJobService(jobRepository)
	JobHandler := job_handler.NewJobHandler(JobService)

	jobRouter := router.Group("/jobs")
	{
		jobRouter.POST("/", JobHandler.CreateJob)
	}
}
