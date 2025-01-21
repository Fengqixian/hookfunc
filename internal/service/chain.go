package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"hookfunc/internal/model"
	"hookfunc/internal/okx"
	"hookfunc/internal/repository"
	"hookfunc/pkg/log"
	"strconv"
	"time"
)

type ChainService interface {
	ChainTransaction(ctx context.Context)
}

func NewChainService(log *log.Logger, rdb *redis.Client, transactionRepository repository.TransactionRepository, config *okx.Config, userInfoRepository repository.UserInfoRepository) ChainService {
	http := okx.NewHttp(config)
	return &chainService{
		log:                   log,
		userInfoRepository:    userInfoRepository,
		Config:                config,
		Http:                  http,
		rdb:                   rdb,
		transactionRepository: transactionRepository,
	}
}

type chainService struct {
	log                *log.Logger
	userInfoRepository repository.UserInfoRepository
	rdb                *redis.Client
	*okx.Config
	*okx.Http
	transactionRepository repository.TransactionRepository
}

func (c *chainService) ChainTransaction(ctx context.Context) {
	var response okx.TransactionRecordResponse
	for _, wallet := range c.Wallets {
		err := c.Http.Get(fmt.Sprintf("https://api.trongrid.io/v1/accounts/%s/transactions/trc20?only_to=true&only_confirmed=true&min_timestamp=%s", wallet.WalletAddress, GetTimestampOneHourAgo()), &response)
		if err != nil {
			return
		}

		for _, record := range response.Data {
			userInfo, err := c.userInfoRepository.FirstByUserWallet(ctx, record.From)
			if err != nil {
				continue
			}

			parseInt, err := strconv.ParseInt(record.Value, 10, 64)
			if err != nil {
				continue
			}

			r, err := c.transactionRepository.ExistByTransactionId(ctx, record.TransactionId)
			if err == nil && r {
				continue
			}

			// 单位天
			var level int
			if record.From == userInfo.Wallet && parseInt > c.Config.SubscriptionPrice[0] {
				if parseInt >= c.Config.SubscriptionPrice[0] && parseInt < c.Config.SubscriptionPrice[1] {
					// 开通一月
					level = 31
					userInfo.SubscriptionEnd = GetNewDateTime(userInfo.SubscriptionEnd, level)

				} else if parseInt >= c.Config.SubscriptionPrice[1] && parseInt < c.Config.SubscriptionPrice[2] {
					// 开通一季
					level = 93
					userInfo.SubscriptionEnd = GetNewDateTime(userInfo.SubscriptionEnd, level)

				} else if parseInt >= c.Config.SubscriptionPrice[2] {
					// 开通一年
					level = 365
					userInfo.SubscriptionEnd = GetNewDateTime(userInfo.SubscriptionEnd, level)

				}

				err := c.userInfoRepository.UpdateSubscriptionEndTime(ctx, userInfo)
				if err != nil {
					c.log.Logger.Error("【链上交易】用户订阅数据保存失败：", zap.Error(err), zap.Any("transaction", userInfo))
					continue
				}

				transaction := model.Transaction{
					UserID:         userInfo.ID,
					TransactionID:  record.TransactionId,
					From:           record.From,
					To:             record.To,
					Value:          record.Value,
					Level:          level,
					BlockTimestamp: record.BlockTimestamp,
					Type:           record.Type,
				}

				err = c.transactionRepository.SaveTransaction(ctx, &transaction)
				if err != nil {
					c.log.Logger.Error("【链上交易】交易数据保存失败：", zap.Error(err), zap.Any("transaction", transaction))
					continue
				}
				c.log.Logger.Info("【链上交易】订阅成功", zap.Any("transaction", transaction))
			}
		}
	}
}

// GetNewDateTime 根据输入的datetime和天数返回新的datetime
func GetNewDateTime(datetime time.Time, days int) time.Time {
	now := time.Now()
	// 检查datetime是否为零值
	if datetime.IsZero() || datetime.Before(now) {
		// 如果datetime为零值，使用当前系统时间
		datetime = now
	}

	// 使用AddDate方法将时间加上指定的天数
	newTime := datetime.AddDate(0, 0, days)

	return newTime
}

func GetTimestampOneHourAgo() string {
	// 获取当前时间
	currentTime := time.Now()

	// 计算一小时前的时间
	oneHourAgo := currentTime.Add(-2000 * time.Hour)

	// 将时间转换为时间戳（以秒为单位）
	timestamp := oneHourAgo.Unix()

	return strconv.FormatInt(timestamp, 10) + "000"
}
