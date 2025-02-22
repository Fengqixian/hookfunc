// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStrategy = "strategy"

// Strategy 策略信息
type Strategy struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                        // 主键ID
	UserID            int64     `gorm:"column:user_id;not null;comment:用户ID" json:"userId"`                                    // 用户ID
	InstId            string    `gorm:"column:inst_id;not null;comment:币种" json:"instId"`                                      // 币种ID
	StrategyName      string    `gorm:"column:strategy_name;not null;comment:策略名称" json:"strategyName"`                        // 策略名称
	SubscriptionState int32     `gorm:"column:subscription_state;not null;comment:订阅状态: 0 未订阅 1 已订阅" json:"subscriptionState"` // 订阅状态: 0 未订阅 1 已订阅
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`  // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`  // 更新时间
	Deleted           bool      `gorm:"column:deleted;not null;comment:是否删除" json:"deleted"`                                   // 是否删除
}

// TableName Strategy's table name
func (*Strategy) TableName() string {
	return TableNameStrategy
}
