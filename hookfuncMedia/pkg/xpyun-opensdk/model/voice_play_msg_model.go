package model

type VoicePlayMsgRequest struct {
	RestRequest `json:",inline"`
	/**
	 * 打印机编号
	 */
	Sn string `json:"sn"`

	/**
	 * 语音类型，01 固定语音播报  02 TTS智能语音播报  02会单独开放接口
	 */
	VoiceType string `json:"voiceType,omitempty"`
	/**
	 * 语音播放次数，最低要求播放1次
	 * 格式：#100  表示播放100次
	 */
	VoiceTime int `json:"voiceTime,omitempty"`
	/**
	 * 语音播放间隔，单位秒，只有当voiceTime>1时才有效
	 * 格式：#100 如果播放次数大于1，则100s之后播放下一次
	 */
	VoiceInterval int `json:"voiceInterval,omitempty"`

	/**
	 * 语音播报内容
	 * 格式：#n1#n2 在语音ic中的序号 目前支持 0-106取值
	 */
	Content string `json:"content"`
}
