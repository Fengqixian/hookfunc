package service

import (
	"context"
	"fmt"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
	"hookfunc/pkg/helper/uuid"
	"math/rand"
	"time"
)

var (
	VerificationsCodeRedisKey = "sms:verifications:code:%s"
)

type UserInfoService interface {
	GetUserInfo(ctx context.Context, openid string) (*model.UserInfo, error)
	GetUserInfoById(ctx context.Context, userId int64) (*model.UserInfo, error)
	CreateUserInfo(ctx context.Context, userInfo *model.UserInfo) error
	SendNoteVerificationCode(ctx context.Context, phone string) error
	VerificationCode(ctx context.Context, params v1.SendSMSCodeRequest) (string, error)
	ConfirmRechargeRecord(ctx context.Context, userId int64) (bool, error)
	UpdateUserInfo(ctx context.Context, userInfo *model.UserInfo) error
}

func NewUserInfoService(service *Service, repository *repository.Repository, userInfoRepository repository.UserInfoRepository) UserInfoService {

	return &userInfoService{
		Service:            service,
		userInfoRepository: userInfoRepository,
		Repository:         repository,
	}
}

type userInfoService struct {
	*Service
	userInfoRepository repository.UserInfoRepository
	*repository.Repository
}

func (s *userInfoService) UpdateUserInfo(ctx context.Context, userInfo *model.UserInfo) error {
	return s.userInfoRepository.UpdateUserInfo(ctx, userInfo)
}

func (s *userInfoService) ConfirmRechargeRecord(ctx context.Context, userId int64) (bool, error) {

	return false, nil
}

func (s *userInfoService) VerificationCode(ctx context.Context, params v1.SendSMSCodeRequest) (string, error) {
	code := s.Repository.GetRedisClient().Get(ctx, fmt.Sprintf(VerificationsCodeRedisKey, params.PhoneNumber))
	if code == nil {
		return "", fmt.Errorf("验证码已过期")
	}

	if code.Val() != params.Code {
		return "", fmt.Errorf("验证码错误")
	}

	info, err := s.GetUserInfo(ctx, params.PhoneNumber)
	if err != nil {
		return "", err
	}

	token, err := s.jwt.GenToken(info.ID, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userInfoService) SendNoteVerificationCode(ctx context.Context, phone string) error {
	codeKey := fmt.Sprintf(VerificationsCodeRedisKey, phone)
	cacheCode := s.Repository.GetRedisClient().Get(ctx, codeKey)
	if cacheCode.Val() != "" {
		return fmt.Errorf("验证码已发送")
	}

	code := uuid.GenerateSMSCode()
	s.Repository.GetRedisClient().Set(ctx, codeKey, code, 120*time.Second)
	return nil
}

var vegetables = []string{
	"番茄", "胡萝卜", "黄瓜", "青椒", "西兰花",
	"菠菜", "生菜", "南瓜", "白萝卜", "卷心菜",
	"茄子", "芹菜", "冬瓜", "小葱", "大蒜",
	"姜", "香菜", "菜花", "苦瓜", "竹笋",
	"豆角", "豌豆", "玉米", "花菜", "蘑菇",
	"洋葱", "韭菜", "辣椒", "大白菜", "红薯",
	"土豆", "青菜", "油菜", "紫甘蓝", "莴笋",
}

// 后缀汉字列表
var suffixes = []string{"君", "酱", "萌"}

// 生成随机昵称的函数
func generateNickname() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	// 随机选择一个蔬菜名称
	vegetable := vegetables[rng.Intn(len(vegetables))]

	// 随机选择一个后缀汉字
	suffix := suffixes[rng.Intn(len(suffixes))]

	// 生成一个随机数作为后缀
	number := rng.Intn(100) // 生成 0-99 的随机数

	// 组合生成昵称
	return fmt.Sprintf("%s%s_%d", vegetable, suffix, number)
}

func (s *userInfoService) GetUserInfoById(ctx context.Context, userId int64) (*model.UserInfo, error) {
	userInfo, err := s.userInfoRepository.FirstByUserId(ctx, userId)
	if err == nil {
		return userInfo, err
	}
	return userInfo, nil
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
		Openid:      openid,
		NickName:    generateNickname(),
		Avatar:      "https://cdn.learnku.com/uploads/images/201805/11/1/ZnrA2VK0SN.png!/both/200x200",
		PhoneNumber: openid,
	}

	err = s.CreateUserInfo(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
