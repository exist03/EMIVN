package service

import (
	"emivn/internal/models"
)

func (s *Service) ApplicationCreate(creater, cardID string, sum int) error {
	err := s.Repo.ApplicationInsert(creater, cardID, sum)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) ApplicationsGetActive() ([]models.Application, error) {
	applications, err := s.Repo.ApplicationGetAll()
	if err != nil {
		return nil, err
	}
	active := make([]models.Application, 0)
	for _, application := range applications {
		if application.Status == true {
			active = append(active, application)
		}
	}
	return active, nil
}
func (s *Service) ApplicationsGetDisputable() ([]models.Application, error) {
	applications, err := s.Repo.ApplicationGetAll()
	if err != nil {
		return nil, err
	}
	disputable := make([]models.Application, 0)
	for _, application := range applications {
		if application.Status == false {
			disputable = append(disputable, application)
		}
	}
	return disputable, nil
}
