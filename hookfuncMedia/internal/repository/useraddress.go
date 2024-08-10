package repository

import (
	"context"
	"hookfunc-media/internal/model"
)

type UserAddressRepository interface {
	GetUserAddress(ctx context.Context, id int64) (*model.UserAddressInfo, error)
	Insert(ctx context.Context, address *model.UserAddressInfo) (*model.UserAddressInfo, error)
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
