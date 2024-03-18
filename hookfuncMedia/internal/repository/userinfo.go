package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
)

type UserInfoRepository interface {
	GetOpenId(jsCode string) (string, error)
	FirstByOpenId(ctx context.Context, openid string) (*model.UserInfo, error)
}

func NewUserInfoRepository(repository *Repository) UserInfoRepository {
	return &userInfoRepository{
		Repository: repository,
	}
}

type userInfoRepository struct {
	*Repository
}

func (r *userInfoRepository) GetOpenId(jsCode string) (string, error) {
	res, err := r.miniProgram.GetAuth().Code2Session(jsCode)
	return res.OpenID, err
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
