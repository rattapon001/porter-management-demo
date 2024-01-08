package app

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

type JobServicePort interface {
	CreateJob(location domain.Location, patient domain.Patient) (*domain.Job, error)
	FindAll() ([]*domain.Job, error)
	FindById(id domain.JobId) (*domain.Job, error)
}

type JobService struct {
	repo domain.JobRepository
}

func NewJobService(repo domain.JobRepository) *JobService {
	return &JobService{
		repo: repo,
	}
}
