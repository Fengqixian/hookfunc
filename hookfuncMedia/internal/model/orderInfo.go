// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrderInfo = "order_info"

// OrderInfo 订单信息
type OrderInfo struct {
	ID           int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                        // 主键ID
	UserID       int64      `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                   // 用户ID
	AddressID    int64      `gorm:"column:address_id;not null;comment:配送地址ID" json:"address_id"`                           // 配送地址ID
	TicketState  int32      `gorm:"column:ticket_state;not null;comment:小票打印状态" json:"ticket_state"`                       // 小票打印状态
	DeliveryTime *time.Time `gorm:"column:delivery_time;comment:配送时间" json:"delivery_time"`                                // 配送时间
	Remark       *string    `gorm:"column:remark;comment:订单备注" json:"remark"`                                              // 订单备注
	CreateTime   time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	Deleted      int32      `gorm:"column:deleted;not null;comment:是否删除" json:"deleted"`                                   // 是否删除
}

// TableName OrderInfo's table name
func (*OrderInfo) TableName() string {
	return TableNameOrderInfo
}
