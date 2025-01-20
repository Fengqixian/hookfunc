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
}

func NewJob(
	logger *log.Logger,
	chainService service.ChainService,
) *Job {
	return &Job{
		logger:       logger,
		chainService: chainService,
	}
}

func (j *Job) Start(ctx context.Context) error {
	j.scheduler = gocron.NewScheduler(time.UTC)
	_, err := j.scheduler.CronWithSeconds("0/5 * * * * *").Do(func() {
		start := time.Now()
		j.chainService.ChainTransaction(ctx)
		elapsed := time.Since(start)
		j.logger.Info("【链上交易数据同步】", zap.Any("elapsed", elapsed))
	})
	if err != nil {
		j.logger.Error("【链上交易数据同步】任务执行失败", zap.Error(err))
	}

	j.scheduler.StartBlocking()
	return nil
}
func (j *Job) Stop(ctx context.Context) error {
	return nil
}
