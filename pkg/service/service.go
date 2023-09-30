package service

import (
	"wb_l0"
	"wb_l0/pkg/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetOrder(id int) (wb_l0.Order, error) {
	return s.repo.GetOrderFromCache(id)
}
