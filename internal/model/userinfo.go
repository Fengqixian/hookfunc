// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserInfo = "user_info"

// UserInfo 用户信息
type UserInfo struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                       // 主键ID
	Openid      string    `gorm:"column:openid;not null;comment:微信openId" json:"openid"`                                // 微信openId
	NickName    string    `gorm:"column:nick_name;not null;comment:用户昵称" json:"nickName"`                               // 用户昵称
	Avatar      string    `gorm:"column:avatar;not null;comment:用户头像" json:"avatar"`                                    // 用户头像
	PhoneNumber *string   `gorm:"column:phone_number;comment:电话" json:"phoneNumber"`                                    // 电话
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
	Deleted     bool      `gorm:"column:deleted;not null;comment:是否删除" json:"deleted"`
}

// TableName UserInfo's table name
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}