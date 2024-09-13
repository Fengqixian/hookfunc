package model

// 打印接口请求参数
type PrintRequest struct {
	RestRequest `json:",inline"`
	/**
	 * 打印机编号
	 */
	Sn string `json:"sn"`

	/**
	 * 打印内容,不能超过12k字节
	 */
	Content string `json:"content"`

	/**
	 * 打印份数，默认为1
	 */
	Copies int `json:"copies,omitempty"`

	/**
	 * 打印模式，默认为0
	 */
	Mode int `json:"mode,omitempty"`

	/**
	 * 支付方式41~55：支付宝 微信 ...
	 */
	PayType int `json:"payType,omitempty"`
	/**
	 * 支付与否59~61：退款 到账 消费
	 */
	PayMode int `json:"payMode,omitempty"`
	/**
	 * 支付金额
	 */
	Money float64 `json:"money,omitempty"`
	/**
	 * 声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，默认为 2 来单播放模式，
	 * 3 播报退单语音，4 播报美团来单语音，5 播报饿了么来单语音
	 */
	Voice int `json:"voice,omitempty"`
	/**
	 * 幂等能力，默认不提供外界使用
	 */
	Idempotent string `json:"idempotent,omitempty"`

	/**
	 * 订单有效期 单位：秒  取值范围为：0<订单失效时间<24*60*60
	 */
	ExpiresIn int `json:"expiresIn,omitempty"`

	/**
	 * 是否支持切刀自定义控制  0采用默认切刀  1采用自定义控制切刀
	 */
	Cutter int `json:"cutter,omitempty"`

	/**
	 * 回调地址对应标识，必须先在管理后台设置，否则无效
	 */
	BackurlFlag int `json:"backurlFlag,omitempty"`
	/**
	 * 附加值 字符长度不能超出100
	 */
	Attached string `json:"attached,omitempty"`

	/**
	 * 语音类型，01 固定语音播报  02 TTS智能语音播报
	 */
	TtsVoiceType string `json:"ttsVoiceType,omitempty"`
	/**
	 * 语音播放次数，最低要求播放1次
	 * 格式：#100  表示播放100次
	 */
	TtsVoiceTime int `json:"ttsVoiceTime,omitempty"`
	/**
	 * 语音播放间隔，单位秒，只有当voiceTime>1时才有效
	 * 格式：#100 如果播放次数大于1，则100s之后播放下一次
	 */
	TtsVoiceInterval int `json:"ttsVoiceInterval,omitempty"`
	/**
	 * 支持来单播放tts语音文本，需要设备支持
	 */
	Tts string `json:"tts,omitempty"`
}
