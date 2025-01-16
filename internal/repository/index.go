package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hookfunc/internal/model"
)

type IndexRepository interface {
	ListIndex(ctx context.Context) (*[]model.Index, error)
	GetIndex(ctx context.Context, id int64) (*model.Index, error)
}

func NewIndexRepository(repository *Repository) IndexRepository {
	return &indexRepository{
		Repository: repository,
	}
}

type indexRepository struct {
	*Repository
}

func (i indexRepository) GetIndex(ctx context.Context, id int64) (*model.Index, error) {
	var index model.Index
	if err := i.DB(ctx).Where("deleted = 0 and id =?", id).Find(&index).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &index, nil
		}
		return nil, err
	}

	return &index, nil
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
