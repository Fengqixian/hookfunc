package service

import (
	"context"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type BarService interface {
	ListBar(ctx context.Context) (*[]model.Bar, error)
	ListCoin(ctx context.Context) (*[]model.Coin, error)
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

func (b barService) ListCoin(ctx context.Context) (*[]model.Coin, error) {
	return b.barRepository.ListCoin(ctx)
}

func (b barService) ListBar(ctx context.Context) (*[]model.Bar, error) {
	return b.barRepository.ListBar(ctx)
}
