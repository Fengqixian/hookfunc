package service

import (
	"hookfunc-media/internal/repository"
	"hookfunc-media/pkg/helper/sid"
	"hookfunc-media/pkg/jwt"
	"hookfunc-media/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
