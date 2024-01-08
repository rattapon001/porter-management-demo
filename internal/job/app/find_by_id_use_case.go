package app

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

func (s *JobService) FindById(id domain.JobId) (*domain.Job, error) {
	return s.repo.FindById(id)
}
