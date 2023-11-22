package service

import (
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type ResourceService interface {
	GetResource(id int64) (*model.Resource, error)
}

func NewResourceService(service *Service, resourceRepository repository.ResourceRepository) ResourceService {
	return &resourceService{
		Service:        service,
		resourceRepository: resourceRepository,
	}
}

type resourceService struct {
	*Service
	resourceRepository repository.ResourceRepository
}

func (s *resourceService) GetResource(id int64) (*model.Resource, error) {
	return s.resourceRepository.FirstById(id)
}
