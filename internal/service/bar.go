package service

import (
	"context"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type BarService interface {
	ListBar(ctx context.Context) (*[]model.Bar, error)
}

func NewBarService(service *Service, barRepository repository.BarRepository) BarService {
	return &barService{
		Service:       service,
		barRepository: barRepository,
	}
}

type barService struct {
	*Service
	barRepository repository.BarRepository
}

func (b barService) ListBar(ctx context.Context) (*[]model.Bar, error) {
	return b.barRepository.ListBar(ctx)
}
