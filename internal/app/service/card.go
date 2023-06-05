package service

import "emivn/internal/models"

func (s *Service) CardInDispute(cardID string) bool {
	return s.Repo.CardInDispute(cardID)
}
func (s *Service) CardGetByID(cardID string) (models.Card, error) {
	return s.Repo.CardGetByID(cardID)
}
func (s *Service) CardSetDisputeTrue(cardID string) error {
	return s.Repo.CardSetDisputeTrue(cardID)
}
