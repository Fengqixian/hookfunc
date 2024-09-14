package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
)

type UserInfoRepository interface {
	FirstByOpenId(ctx context.Context, openid string) (*model.UserInfo, error)
	FirstByUserId(ctx context.Context, userId int64) (*model.UserInfo, error)
	InsertUserInfo(ctx context.Context, userInfo *model.UserInfo) error
}

func NewUserInfoRepository(repository *Repository) UserInfoRepository {
	return &userInfoRepository{
		Repository: repository,
	}
}

type userInfoRepository struct {
	*Repository
}

func (r *userInfoRepository) FirstByUserId(ctx context.Context, userId int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	if err := r.DB(ctx).Where("id = ?", userId).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &userInfo, nil
}

func (r *userInfoRepository) InsertUserInfo(ctx context.Context, userInfo *model.UserInfo) error {
	if err := r.DB(ctx).Create(userInfo).Error; err != nil {
		return err
	}
	return nil
}

func (r *userInfoRepository) FirstByOpenId(ctx context.Context, id string) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	if err := r.DB(ctx).Where("openid = ?", id).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &userInfo, nil
}
