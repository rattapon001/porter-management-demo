package msg_relay

import (
	"context"
	"fmt"
	"log"

	"github.com/rattapon001/porter-management-demo/internal/job/domain"
	job_mongo "github.com/rattapon001/porter-management-demo/internal/job/infra/mongo"
)

func MsgRelay(ctx context.Context) {
	db := job_mongo.MongoDbInit()
	changeStream := ChangeCapture(db, "porter_management", "jobs", ctx)
	for changeStream.Next(ctx) {
		var changeEvent struct {
			FullDocument *domain.Job `bson:"fullDocument"`
		}
		if err := changeStream.Decode(&changeEvent); err != nil {
			log.Println("Error decoding change event:", err)
			continue
		}
		// Handle the change event as needed
		if changeEvent.FullDocument != nil {
			// Handle the change event as needed
			fmt.Printf("Job: %+v\n", *changeEvent.FullDocument)
		}
	}
}
