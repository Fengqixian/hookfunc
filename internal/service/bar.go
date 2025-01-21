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
		cache:         make(map[string]interface{}),
	}
}

type barService struct {
	*Service
	barRepository repository.BarRepository
	cache         map[string]interface{}
}

func (b barService) ListCoin(ctx context.Context) (*[]model.Coin, error) {
	if b.cache["coins"] != nil {
		return b.cache["coins"].(*[]model.Coin), nil
	}

	result, err := b.barRepository.ListCoin(ctx)
	if err != nil {
		return nil, err
	}

	b.cache["coins"] = result
	return result, nil
}

func (b barService) ListBar(ctx context.Context) (*[]model.Bar, error) {
	if b.cache["bars"] != nil {
		return b.cache["bars"].(*[]model.Bar), nil
	}
	result, err := b.barRepository.ListBar(ctx)
	if err != nil {
		return nil, err
	}
	b.cache["bars"] = result
	return result, nil
}
