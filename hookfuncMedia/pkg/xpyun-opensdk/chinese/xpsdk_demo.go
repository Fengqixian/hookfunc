package chinese

import (
	"encoding/json"
	"fmt"
	"xpyun-opensdk/model"
	"xpyun-opensdk/service"
	"xpyun-opensdk/util"
)

/*
*
*必填*：开发者ID：芯烨云后台注册账号（即邮箱地址或开发者ID），开发者用户注册成功之后，登录芯烨云后台，在【个人中心=》账号信息】下可查看开发者ID

当前【XXXXXXXXXXXXXX】只是样例，需修改再使用
*/
const USER_NAME = "570311238@qq.com"

/*
*
*必填*：开发者密钥：芯烨云后台注册账号后自动生成的开发者密钥，开发者用户注册成功之后，登录芯烨云后台，在【个人中心=》账号信息】下可查看开发者密钥

当前【XXXXXXXXXXXXXX】只是样例，需修改再使用
*/
const USER_KEY = "ef1d7799f8704d16badac4a0ae5aa94a"

/*
*
*必填*：打印机设备编号，必须要在芯烨云管理后台的【打印管理->打印机管理】下添加打印机或调用API接口添加打印机，测试小票机和标签机的时候注意替换打印机编号
打印机设备编号获取方式：在打印机底部会有带PID或SN字样的二维码且PID或SN后面的一串字符即为打印机编号

当前【XXXXXXXXXXXXXX】只是样例，需修改再使用
*/
const OK_PRINTER_SN = "00RJMSDUBQ15E4A"

// 1.批量添加打印机
func AddPrintersTest() {
	request := model.AddPrinterRequest{}

	requestItem1 := model.AddPrinterRequestItem{}
	requestItem1.Sn = OK_PRINTER_SN
	// 如果名称不指定，将默认为 芯烨云开放平台打印机
	// requestItem1.Name = "测试打印机"

	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()

	request.Items = make([]*model.AddPrinterRequestItem, 0)
	request.Items = append(request.Items, &requestItem1)

	result := service.XpYunAddPrinters(&request)

	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

/**
 * 2.设置打印机语音类型
 * @param restRequest
 * @return
 */
func SetVoiceTypeTest() {
	request := model.SetVoiceTypeRequest{}

	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()

	request.Sn = OK_PRINTER_SN

	// 声音类型：
	// 打印机固件版本为V10.xx的机器取值： 0真人语音（大） 1真人语音（中） 2真人语音（小） 3 嘀嘀声 4 静音
	// 其它固件版本的机器取值：0真人语音 3 嘀嘀声 4 静音
	request.VoiceType = 2
	// 声音大小：0大 1中 2小 3关闭
	// 说明：打印机固件版本为非V10.xx的机器支持此参数
	request.VolumeLevel = 0

	result := service.XpYunSetVoiceType(request)

	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 3.批量删除打印机
func DelPrintersTest() {
	request := model.DelPrinterRequest{}

	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()

	snList := []string{OK_PRINTER_SN}
	request.SnList = snList

	result := service.XpYunDelPrinters(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 4.修改打印机信息
func UpdPrinterTest() {
	request := model.UpdPrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()
	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN

	//*必填*：打印机名称
	request.Name = "X58C111"

	result := service.XpYunUpdatePrinter(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 5.清空待打印队列
func XpYunDelPrinterQueueTest() {
	request := model.ClearPrintOrderRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN
	// 可以根据指定订单号清除某一单打印
	// 打印接口返回订单号，以OM开头
	// request.OrderId = "OM21110714211742658474"
	request.GenerateSign()

	result := service.XpYunDelPrinterQueue(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 6.查询订单是否打印成功
func XpYunQueryOrderStateTest() {
	request := model.QueryOrderStateRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()
	request.GenerateSign()

	// *必填*：订单编号，由“打印订单”接口返回 以OM开头
	request.OrderId = "OM20110715072844111475"

	result := service.XpYunQueryOrderState(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 7.查询打印机某天的订单统计数
func XpYunQueryOrderStatisTest() {
	request := model.QueryOrderStatisRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN

	//*必填*：查询日期，格式yyyy-MM-dd，如：2019-08-15
	request.Date = "2020-10-03"

	result := service.XpYunQueryOrderStatis(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 8.查询打印机状态
func XpYunQueryPrinterStatusTest() {
	request := model.PrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	result := service.XpYunQueryPrinterStatus(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 8.批量查询打印机状态
func XpYunQueryPrintersStatusTest() {
	request := model.PrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	result := service.XpYunQueryPrinterStatus(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 9.金额播报
func XpYunPlayVoiceTest() {
	request := model.VoiceRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	//支付方式：
	//取值范围41~55：
	//支付宝 41、微信 42、云支付 43、银联刷卡 44、银联支付 45、会员卡消费 46、会员卡充值 47、翼支付 48、成功收款 49、嘉联支付 50、壹钱包 51、京东支付 52、快钱支付 53、威支付 54、享钱支付 55
	//仅用于支持金额播报的芯烨云打印机。
	request.PayType = 41

	//支付与否：
	//取值范围59~61：
	//退款 59 到账 60 消费 61。
	//仅用于支持金额播报的芯烨云打印机
	request.PayMode = 59

	//支付金额：
	//最多允许保留2位小数。
	//仅用于支持金额播报的芯烨云打印机。
	request.Money = 24.15

	result := service.XpYunPlayVoice(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 14、钱箱控制
func XpYunControlBoxTest() {
	request := model.PrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	result := service.XpYunControlBox(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 15、扩展语音播报
// 发送用户需要播报的语音内容给支持预定义语音的芯烨云打印机
func XpYunPlayVoiceExtTest() {
	request := model.VoicePlayMsgRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()
	// 常用的语音播报内容
	// 语音播报内容，如#62，将播报序号为 62 对应的语音内容。
	// 该内容需加入芯烨云技术支持 QQ 群（856926694）咨询
	/**
	5 钱箱打开
	8 网络断开，请检查！
	13 您有新的订单！
	14 小票打印完成
	21 您有新的美团外卖订单，请及时处理
	22 您有新的饿了么订单，请及时处理
	23 您有新的订单，请及时处理
	24 有用户申请取消订单了
	25 有用户申请退单了
	26 嘀声
	27 嘀声
	28 嘀嘀嘀声
	92 顾客请取餐
	93 顾客请就餐
	94 顾客加菜
	95 顾客退菜
	62 0
	63 1
	64 2
	65 3
	66 4
	67 5
	68 6
	69 7
	70 8
	71 9
	72 点
	73 元
	74 十
	75 百
	76 千
	77 万
	*/
	request.Content = "#13#23#24"
	//语音播报次数，默认为 1 次
	request.VoiceTime = 1
	// 语音播报多次，当前次播报与下一次播报的时间间隔，只有
	// 当播报次数大于 1 时才有效。
	request.VoiceInterval = 1

	result := service.XpYunPlayVoiceExt(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 16、自定义语音播报
// 发送用户需要播报的语音内容给支持自定义语音播报的芯烨云打印机，若有需要可联系技
// 术或商务购买支持自定义语音播报的机型
func XpYunPlayCustomVoiceTest() {
	request := model.VoicePlayMsgRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()
	request.Content = "A9十2 号顾客请取餐"
	//语音播报次数，默认为 1 次
	request.VoiceTime = 1
	// 语音播报多次，当前次播报与下一次播报的时间间隔，只有
	// 当播报次数大于 1 时才有效。
	request.VoiceInterval = 1

	result := service.XpYunPlayCustomVoice(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 17、店铺 LOGO 上传
// 开发者可以使用 LOGO 上传接口，上传商户店铺 LOGO，标签打印机在小票模式下不支持 LOGO
// 上传打印，可以在订单内容里面使用<IMG200>LOGO 的 base64 内容</IMG>这个方式实现 LOGO
// 打印。在实际使用该接口时，打印机需要出于空闲状态且已经开机并完成了配网操
// 作。LOGO 执行完毕后打印机有播音“您有新的订单，请及时处理”提示
func XpYunUploadLogoTest() {
	request := model.UploadLogoRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	// 店铺 LOGO 的 base64 格式内容，不包含图片格式的前缀符。
	// 具体参见下发样例，content 需要要使用真实数据
	request.Content = "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"
	//标签机（T271U、320B）有效且需要指定固定值 2
	// request.LabelMode = 2
	// LOGO 的图片大小，标签机有效，指定该值时，LOGO 图片的
	// 大小宽或高比该值大时，将以此值为基准进行等比缩放。未
	// 达到该值时，保留原图大小不作处理
	// request.ImageSize = 200

	result := service.XpYunUploadLogo(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 18、店铺 LOGO 删除
func XpYunDelUploadLogoTest() {
	request := model.PrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	result := service.XpYunDelUploadLogo(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 19、获取打印机基本信息
func XpYunPrinterInfoTest() {
	request := model.PrinterRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY
	request.Sn = OK_PRINTER_SN
	request.GenerateSign()

	result := service.XpYunPrinterInfo(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}
