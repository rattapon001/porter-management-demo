package job_mongo

import (
	"context"

	"github.com/rattapon001/porter-management-demo/internal/job/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobMongoRepository struct {
	Collection *mongo.Collection
}

func NewJobMongoRepository(collection *mongo.Collection) *JobMongoRepository {
	return &JobMongoRepository{
		Collection: collection,
	}
}

func (r *JobMongoRepository) Save(job *domain.Job) error {
	_, err := r.Collection.InsertOne(context.Background(), job)
	return err
}

func (r *JobMongoRepository) Update(job *domain.Job) error {
	_, err := r.Collection.UpdateOne(context.Background(), domain.Job{Id: job.Id}, job)
	return err
}

func (r *JobMongoRepository) FindById(id domain.JobId) (*domain.Job, error) {
	var job domain.Job
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&job)
	return &job, err
}

func (r *JobMongoRepository) FindAll() ([]*domain.Job, error) {
	filter := bson.D{}
	var jobs []*domain.Job
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var job domain.Job
		err := cursor.Decode(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

func (r *JobMongoRepository) Delete(job *domain.Job) error {
	_, err := r.Collection.DeleteOne(context.Background(), domain.Job{Id: job.Id})
	return err
}
