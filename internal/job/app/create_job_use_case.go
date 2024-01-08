package app

import (
	"github.com/rattapon001/porter-management-demo/internal/job/domain"
)

func (s *JobService) CreateJob(location domain.Location, patient domain.Patient, name string) (*domain.Job, error) {
	job := domain.CreateNewJob(location, patient, name)
	eventData := domain.EventDataCreated{
		ID:       job.ID,
		Location: job.Location,
		Patient:  job.Patient,
		Version:  job.Version,
		Name:     job.Name,
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
