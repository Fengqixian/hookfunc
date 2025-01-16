package service

import (
	"context"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type IndexService interface {
	ListIndex(ctx context.Context) (*[]model.Index, error)
}

func NewIndexService(service *Service, indexRepository repository.IndexRepository) IndexService {
	return &indexService{
		Service:         service,
		indexRepository: indexRepository,
	}
}

type indexService struct {
	*Service
	indexRepository repository.IndexRepository
}

func (i indexService) ListIndex(ctx context.Context) (*[]model.Index, error) {
	return i.indexRepository.ListIndex(ctx)
}
