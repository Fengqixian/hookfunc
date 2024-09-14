package v1

type UserAddressInfoRequest struct {
	UserId        int64  `json:"userId"`
	Province      string `json:"province" binding:"required"`      // 省
	City          string `json:"city" binding:"required"`          // 市
	Region        string `json:"region" binding:"required"`        // 区
	Longitude     int64  `json:"longitude" binding:"required"`     // 经度
	Latitude      int64  `json:"latitude" binding:"required"`      // 纬度
	UserName      string `json:"userName" binding:"required"`      // 收获人姓名
	PhoneNumber   string `json:"phoneNumber" binding:"required"`   // 收获人电话
	AddressDetail string `json:"addressDetail" binding:"required"` // 详细地址
}
