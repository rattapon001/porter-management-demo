package app

import (
	"github.com/rattapon001/porter-management-demo/internal/job/domain"
)

func (s *JobService) CreateJob(location domain.Location, patient domain.Patient) (*domain.Job, error) {
	job := domain.CreateNewJob(location, patient)
	eventData := domain.EventDataCreated{
		ID:       job.ID,
		Location: job.Location,
		Patient:  job.Patient,
		Version:  job.Version,
	}
	event := domain.Event{
		Type: "Job." + domain.JobCreated,
		Data: eventData,
	}
	job.AppendEvent(event)
	err := s.repo.Save(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
