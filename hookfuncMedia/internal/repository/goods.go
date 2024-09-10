package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
)

type GoodsRepository interface {
	GetGoods(ctx context.Context, id int64) (*model.GoodsInfo, error)
	GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error)
	ListGoodsById(ctx context.Context, ids []int64) (*[]model.GoodsInfo, error)
}

func NewGoodsRepository(
	repository *Repository,
) GoodsRepository {
	return &goodsRepository{
		Repository: repository,
	}
}

type goodsRepository struct {
	*Repository
}

func (r *goodsRepository) ListGoodsById(ctx context.Context, ids []int64) (*[]model.GoodsInfo, error) {
	var goods []model.GoodsInfo
	if err := r.DB(ctx).Where("deleted = 0 and publish_state = 1 and id in ?", ids).Find(&goods).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &goods, nil
}

func (r *goodsRepository) GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error) {
	var goods []model.GoodsInfo
	if err := r.DB(ctx).Where("deleted = 0 and publish_state = 1").Order("sort desc").Find(&goods).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &goods, nil
}

func (r *goodsRepository) GetGoods(ctx context.Context, id int64) (*model.GoodsInfo, error) {
	var goods model.GoodsInfo
	if err := r.DB(ctx).Where("deleted = 0 and publish_state = 1 and id = ?", id).First(&goods).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &goods, nil
}
