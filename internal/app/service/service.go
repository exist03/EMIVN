package service

import (
	"emivn/internal/app/repository"
	"strconv"
	"time"
)

type Service struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) SamuraiInputTurnover(amount, username, bank string) error {
	date := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	sum, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}
	err = s.Repo.SamuraiSetTurnover(username, sum, date, bank)
	if err != nil {
		return err
	}
	return nil
}
