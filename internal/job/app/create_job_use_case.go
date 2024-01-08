package app

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

func (s *JobService) CreateJob(location domain.Location, patient domain.Patient, name string) (*domain.Job, error) {
	job := domain.CreateNewJob(location, patient, name)
	err := s.repo.Save(job)
	if err != nil {
		return nil, err
	}
	// publish event jobCreated
	return job, nil
}
