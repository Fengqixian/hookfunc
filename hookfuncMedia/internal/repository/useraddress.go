package repository

import (
	"context"
	"hookfunc-media/internal/model"
)

type UserAddressRepository interface {
	GetUserAddress(ctx context.Context, id int64) (*model.UserAddressInfo, error)
	Insert(ctx context.Context, address *model.UserAddressInfo) (*model.UserAddressInfo, error)
	GetAddressByIds(ctx context.Context, ids []int64) ([]model.UserAddressInfo, error)
}

func NewUserAddressRepository(
	repository *Repository,
) UserAddressRepository {
	return &userAddressRepository{
		Repository: repository,
	}
}

type userAddressRepository struct {
	*Repository
}

func (r *userAddressRepository) GetAddressByIds(ctx context.Context, ids []int64) ([]model.UserAddressInfo, error) {
	var list []model.UserAddressInfo
	if err := r.DB(ctx).Where("id in ?", ids).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (r *userAddressRepository) Insert(ctx context.Context, address *model.UserAddressInfo) (*model.UserAddressInfo, error) {
	if err := r.DB(ctx).Create(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *userAddressRepository) GetUserAddress(ctx context.Context, id int64) (*model.UserAddressInfo, error) {
	var userAddress model.UserAddressInfo

	return &userAddress, nil
}
