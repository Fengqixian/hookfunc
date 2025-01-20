// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTransaction = "transaction"

// Transaction 交易信息
type Transaction struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                        // 主键ID
	UserID         int64     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                   // 用户ID
	Level          int32     `gorm:"column:level;not null;comment:开通级别" json:"level"`                                       // 开通级别
	TransactionID  string    `gorm:"column:transaction_id;not null;comment:转账记录ID" json:"transaction_id"`                   // 转账记录ID
	BlockTimestamp int64     `gorm:"column:block_timestamp;not null;comment:区块确认时间" json:"block_timestamp"`                 // 区块确认时间
	From           string    `gorm:"column:from;not null;comment:来自某一账号的转账" json:"from"`                                    // 来自某一账号的转账
	To             string    `gorm:"column:to;not null;comment:收款账号" json:"to"`                                             // 收款账号
	Type           string    `gorm:"column:type;not null;comment:类型" json:"type"`                                           // 类型
	Value          string    `gorm:"column:value;not null;comment:数量" json:"value"`                                         // 数量
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	Deleted        bool      `gorm:"column:deleted;not null;comment:是否删除" json:"deleted"`                                   // 是否删除
}

// TableName Transaction's table name
func (*Transaction) TableName() string {
	return TableNameTransaction
}
