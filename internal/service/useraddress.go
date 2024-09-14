package service

import (
	"context"
	"github.com/mitchellh/mapstructure"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type UserAddressService interface {
	UpdateAddress(ctx context.Context, params *model.UserAddressInfo) error
	ListUserAddresses(ctx context.Context, userId int64) (*[]model.UserAddressInfo, error)
	GetUserAddress(ctx context.Context, id int64) (*model.UserAddressInfo, error)
	CreateUserAddress(ctx context.Context, params *v1.UserAddressInfoRequest) (*model.UserAddressInfo, error)
}

func NewUserAddressService(
	service *Service,
	userAddressRepository repository.UserAddressRepository,
) UserAddressService {
	return &userAddressService{
		Service:               service,
		userAddressRepository: userAddressRepository,
	}
}

type userAddressService struct {
	*Service
	userAddressRepository repository.UserAddressRepository
}

func (s *userAddressService) UpdateAddress(ctx context.Context, params *model.UserAddressInfo) error {
	return s.userAddressRepository.Update(ctx, params)
}

func (s *userAddressService) ListUserAddresses(ctx context.Context, userId int64) (*[]model.UserAddressInfo, error) {
	return s.userAddressRepository.ListUserAddresses(ctx, userId)
}

func (s *userAddressService) CreateUserAddress(ctx context.Context, params *v1.UserAddressInfoRequest) (*model.UserAddressInfo, error) {
	var e model.UserAddressInfo
	err := mapstructure.Decode(params, &e)
	entity, err := s.userAddressRepository.Insert(ctx, &e)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *userAddressService) GetUserAddress(ctx context.Context, id int64) (*model.UserAddressInfo, error) {
	return s.userAddressRepository.GetUserAddress(ctx, id)
}
