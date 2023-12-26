package app

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

func (s *JobService) CreateJob(location domain.Location, patient domain.Patient) (*domain.Job, error) {
	job := domain.CreateNewJob(location, patient)
	err := s.repo.Save(job)
	if err != nil {
		return nil, err
	}
	// publish event jobCreated
	return job, nil
}
