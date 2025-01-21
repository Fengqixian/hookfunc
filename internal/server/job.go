package server

import (
	"context"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"hookfunc/internal/service"
	"hookfunc/pkg/log"
	"time"
)

type Job struct {
	logger       *log.Logger
	chainService service.ChainService
	scheduler    *gocron.Scheduler
	lineService  service.LineService
}

func NewJob(
	logger *log.Logger,
	chainService service.ChainService,
	lineService service.LineService,
) *Job {
	return &Job{
		logger:       logger,
		chainService: chainService,
		lineService:  lineService,
	}
}

func (j *Job) Start(ctx context.Context) error {
	j.scheduler = gocron.NewScheduler(time.UTC)
	// 配置钱包收到转账后为钱包所关联的账号开通数据订阅服务
	_, err := j.scheduler.CronWithSeconds("0/5 * * * * *").Do(func() {
		start := time.Now()
		j.chainService.ChainTransaction(ctx)
		elapsed := time.Since(start)
		j.logger.Info("【链上交易数据同步】", zap.Any("elapsed", elapsed))
	})
	if err != nil {
		j.logger.Error("【链上交易数据同步】任务执行失败", zap.Error(err))
	}

	// 更新链上数据
	_, err = j.scheduler.CronWithSeconds("0/1 * * * * *").Do(func() {
		j.lineService.AllCoinChainKlineDataSync(ctx)
	})
	if err != nil {
		j.logger.Error("【OKX数据同步】任务执行失败", zap.Error(err))
	}
	//TODO 为已订阅的用户进行数据推送

	j.scheduler.StartBlocking()
	return nil
}
func (j *Job) Stop(ctx context.Context) error {
	return nil
}
