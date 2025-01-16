package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hookfunc/internal/model"
)

type IndexRepository interface {
	ListIndex(ctx context.Context) (*[]model.Index, error)
}

func NewIndexRepository(repository *Repository) IndexRepository {
	return &indexRepository{
		Repository: repository,
	}
}

type indexRepository struct {
	*Repository
}

func (i indexRepository) ListIndex(ctx context.Context) (*[]model.Index, error) {
	var index []model.Index
	if err := i.DB(ctx).Where("deleted = 0 order by id asc").Find(&index).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &index, nil
		}
		return nil, err
	}

	return &index, nil
}
