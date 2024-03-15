package service

import (
	"context"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type UserInfoService interface {
	GetUserInfo(ctx context.Context, openid string) (*model.UserInfo, error)
}

func NewUserInfoService(service *Service, userInfoRepository repository.UserInfoRepository) UserInfoService {
	return &userInfoService{
		Service:            service,
		userInfoRepository: userInfoRepository,
	}
}

type userInfoService struct {
	*Service
	userInfoRepository repository.UserInfoRepository
}

func (s *userInfoService) GetUserInfo(ctx context.Context, jsCode string) (*model.UserInfo, error) {
	openid, err := s.userInfoRepository.GetOpenId(jsCode)
	if err != nil {
		return nil, err
	}

	return s.userInfoRepository.FirstByOpenId(ctx, openid)
}
