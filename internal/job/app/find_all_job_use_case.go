package app

import "github.com/rattapon001/porter-management-demo/internal/job/domain"

func (s *JobService) FindAll() ([]*domain.Job, error) {
	return s.repo.FindAll()
}
