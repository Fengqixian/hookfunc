package repository

import (
	"hookfunc/internal/model"
)

type ResourceRepository interface {
	FirstById(id int64) (*model.Resource, error)
}

func NewResourceRepository(repository *Repository) ResourceRepository {
	return &resourceRepository{
		Repository: repository,
	}
}

type resourceRepository struct {
	*Repository
}

func (r *resourceRepository) FirstById(id int64) (*model.Resource, error) {
	var resource model.Resource
	// TODO: query db
	return &resource, nil
}
