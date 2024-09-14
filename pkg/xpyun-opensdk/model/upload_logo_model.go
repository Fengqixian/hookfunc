package model

type UploadLogoRequest struct {
	RestRequest `json:",inline"`

	/**
	 * 打印机编号
	 */
	Sn string `json:"sn"`
	// 店铺 LOGO 的 base64 格式内容，不包含图片格式的前缀符。
	// 具体参见下发样例，content 需要要使用真实数据
	Content string `json:"content"`
	//标签机（T271U、320B）有效且需要指定固定值 2
	LabelMode int `json:"labelMode,omitempty"`
	//LOGO 的图片大小，标签机有效，指定该值时，LOGO 图片的
	// 大小宽或高比该值大时，将以此值为基准进行等比缩放。未达到该值时，保留原图大小不作处理
	ImageSize int `json:"imageSize,omitempty"`
}
