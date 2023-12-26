package memory

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

type JobMemoryRepository struct {
	Jobs []*domain.Job
}

func NewJobMemoryRepository() *JobMemoryRepository {
	return &JobMemoryRepository{
		Jobs: []*domain.Job{},
	}
}

func (r *JobMemoryRepository) Save(job *domain.Job) error {
	r.Jobs = append(r.Jobs, job)
	return nil
}

func (r *JobMemoryRepository) Update(job *domain.Job) error {
	for i, j := range r.Jobs {
		if j.Id == job.Id {
			r.Jobs[i] = job
			return nil
		}
	}
	return nil
}

func (r *JobMemoryRepository) FindById(id domain.JobId) (*domain.Job, error) {
	for _, j := range r.Jobs {
		if j.Id == id {
			return j, nil
		}
	}
	return nil, nil
}

func (r *JobMemoryRepository) FindAll() ([]*domain.Job, error) {
	return r.Jobs, nil
}

func (r *JobMemoryRepository) Delete(job *domain.Job) error {
	for i, j := range r.Jobs {
		if j.Id == job.Id {
			r.Jobs = append(r.Jobs[:i], r.Jobs[i+1:]...)
			return nil
		}
	}
	return nil
}
