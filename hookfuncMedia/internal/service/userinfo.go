package service

import (
	"context"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type UserInfoService interface {
	GetUserInfo(ctx context.Context, openid string) (*model.UserInfo, error)
	CreateUserInfo(ctx context.Context, userInfo *model.UserInfo) error
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

func (s *userInfoService) CreateUserInfo(ctx context.Context, userInfo *model.UserInfo) error {
	err := s.userInfoRepository.InsertUserInfo(ctx, userInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *userInfoService) GetUserInfo(ctx context.Context, openid string) (*model.UserInfo, error) {
	userInfo, err := s.userInfoRepository.FirstByOpenId(ctx, openid)
	if err == nil {
		return userInfo, err
	}

	newUser := &model.UserInfo{
		Openid:   openid,
		NickName: "test",
		Avatar:   "https://cdn.learnku.com/uploads/images/201805/11/1/ZnrA2VK0SN.png!/both/200x200",
	}

	err = s.CreateUserInfo(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
