package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hookfunc-media/pkg/log"
	"moul.io/zapgorm2"
	"time"
)

const ctxTxKey = "TxKey"

type Repository struct {
	miniProgram *miniprogram.MiniProgram
	db          *gorm.DB
	rdb         *redis.Client
	logger      *log.Logger
}

func NewRepository(miniProgram *miniprogram.MiniProgram, db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		miniProgram: miniProgram,
		db:          db,
		rdb:         rdb,
		logger:      logger,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// DB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	return db
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}

func NewWechatMiniProgram(conf *viper.Viper) *miniprogram.MiniProgram {
	wc := InitWechat(conf)
	offCfg := &config.Config{
		AppID:     conf.GetString("wechat.appid"),
		AppSecret: conf.GetString("wechat.secret"),
	}
	return wc.GetMiniProgram(offCfg)
}

func InitWechat(conf *viper.Viper) *wechat.Wechat {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        conf.GetString("data.redis.addr"),
		Password:    conf.GetString("data.redis.password"),
		Database:    conf.GetInt("data.redis.db"),
		MaxActive:   100,
		MaxIdle:     100,
		IdleTimeout: 30000,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	return wc
}

func (r *Repository) GetWechatMiniProgram() *miniprogram.MiniProgram {
	return r.miniProgram
}
