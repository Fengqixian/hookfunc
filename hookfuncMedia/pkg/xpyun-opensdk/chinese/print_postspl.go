package chinese

//esc/pos 指令打印票据  tspl指令打印标签使用样例
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"xpyun-opensdk/model"
	"xpyun-opensdk/service"
	"xpyun-opensdk/util"
)

/*
*
POS 指令打印  若对指令不熟悉的开发者,请慎用.在使用非打印内容指令时,需将mode设置为0,否则会出现一些未知的现象

	将用户需要打印的订单内容采用 pos 指令编码后，使用 base64 加密发给芯烨云小票打印
	机。Pos 指令编程文档参见 https://www.xprinter.net/companyfile/1/ 58 打印机参考
	《芯烨 58 系列中文编程手册》或 80 打印机参考《芯烨 80 系列中文编程手册》
*/
func XpYunPosTest() {
	//58 pos指令综合排版base64数据
	printContent := "G2EBHSER0L7sx9TG0KHGsR0hAA0KDQobYQCyy8P7ICAgICAgICAgICAgICAgIMr9wb8gILWlvNsgIA0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCr/JwNa8prPhICAgICAgICAgICAgMiAgICAgOS45OSAgDQrLrtbz0+PM2MCxICAgICAgICAgIDEgICAgIDEwOC4wMA0KusC7qrDms6y8ts7etdDB+s+6s7S3uQ0KICAgICAgICAgICAgICAgICAgICAxICAgICA5OS45MCANCsy/v773qNPjICAgICAgICAgICAgNSAgICAgMTkuOTkgDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KG2ECus+8xqO6MzI3Ljgz1KoNChthAA0KG2EAv827p7XY1rejutbpuqPK0M/j1t7H+Hh4wrd4eLrFDQq/zbuntee7sKO6MTM2MyoqKioqODgNCs/CtaXKsbzko7oyMDIwLTktOSAxNTowNzo1Nw0KsbjXoqO6ydm3xcCxILK7s9TP47LLDQobYQEbMwAbYTEdKGsDADFDBh0oawMAMUUwHShrGAAxUDBodHRwczovL3d3dy54cHl1bi5uZXQdKGsDADFRMBsyHUwAABthAA0KDQo="

	request := model.PrintRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN

	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = printContent

	//打印份数，默认为1
	request.Copies = 1

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
	request.Voice = 2

	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1
	// 支持来单播放tts语音文本，目前不能播英语，若英语单词，将会按照单字母方式逐个播报英文字母，需要设备支持
	// request.Tts = "芯烨云来单了，请及时处理"
	result := service.XpYunPos(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

func XpYunPosTest80() {
	//80 pos指令综合排版base64数据
	printContent := "G2EBHSER0L7sx9TGODDQocax19u6z8XFsOYdIQANCg0KG2EAssvD+yAgICAgICAgICAgICAgICAgICAgICAgIMr9wb8gICC1pbzbICAg19y82yAgDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCr/JwNa8prPhusC7qrDms6y8ts7etdDB+iAgIDIwICAgICA5Ljk1ICAgMTk5LjAwIA0Kz7qztLe5usC7qrDms6y8ts7etdDB+s+6DQqztLe5usC7qrDms6y8ts7etdDB+s+6s7QNCre5DQrLrtbz0+PM2MCxy67W89PjzNjAsSAgICAgICAxICAgICAgMTA4LjAwIDEwOC4wMCANCrrAu6qw5rOsvLbO3rXQwfrPurO0t7nB+iAgIDEgICAgICA5OS45MCAgOTkuOTAgIA0Kz7qztLe5DQrMv7++96jT4yAgICAgICAgICAgICAgICAgICA1ICAgICAgMTkuOTkgIDk5Ljk1ICANCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KG2ECus+8xqO6MzI3Ljc11KoNChthAA0KG2EAv827p7XY1rejutbpuqPK0M/j1t7H+Hh4wrd4eLrFDQq/zbuntee7sKO6MTM2MyoqKioqODgNCs/CtaXKsbzko7oyMDIwLTktOSAxNTowNzo1Nw0KsbjXoqO6ydm3xcCxILK7s9TP47LLDQobYQEbMwAbYTEdKGsEADFBMAAdKGsDADFDBh0oawMAMUUwHShrGAAxUDBodHRwczovL3d3dy54cHl1bi5uZXQdKGsDADFRMBsyHUwAABthAA0KDQo="

	request := model.PrintRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN

	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = printContent

	//打印份数，默认为1
	request.Copies = 1

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
	request.Voice = 2

	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1
	result := service.XpYunPos(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

// 58 T271U标签打印机打印指令
func XpYunPosTsplTest() {
	//80 pos指令综合排版base64数据 末尾一定要换行 保留一个空行
	printContent := `SIZE 40 mm, 30 mm
DIRECTION 0,0
GAP 2.0 mm,0.0 mm
SHIFT 0
DENSITY 7
CLS
TEXT 312,232 ,"TSS24.BF2" ,180 ,1 ,1 ,"#001    一号桌    1/3"
TEXT 312,144 ,"TSS24.BF2" ,180 ,2 ,2 ,"黄金炒饭"
TEXT 312,40 ,"TSS24.BF2" ,180 ,1 ,1 ,"王女士    136****3388"
PRINT 1, 1
SIZE 40 mm, 30 mm
DIRECTION 0,0
GAP 2.0 mm,0.0 mm
SHIFT 0
DENSITY 7
CLS
TEXT 312,232 ,"TSS24.BF2" ,180 ,1 ,1 ,"#001    一号桌    2/3"
TEXT 312,144 ,"TSS24.BF2" ,180 ,2 ,2 ,"凉拌青瓜"
TEXT 312,40 ,"TSS24.BF2" ,180 ,1 ,1 ,"王女士    136****3388"
PRINT 1, 1
SIZE 40 mm, 30 mm
DIRECTION 0,0
GAP 2.0 mm,0.0 mm
SHIFT 0
DENSITY 7
CLS
TEXT 312,232 ,"TSS24.BF2" ,180 ,1 ,1 ,"#001    一号桌    3/3"
TEXT 312,144 ,"TSS24.BF2" ,180 ,2 ,2 ,"老刘家肉夹馍"
TEXT 312,40 ,"TSS24.BF2" ,180 ,1 ,1 ,"王女士    136****3388"
PRINT 1, 1
SIZE 40 mm, 30 mm
DIRECTION 0,0
GAP 2.0 mm,0.0 mm
SHIFT 0
DENSITY 7
CLS
TEXT 312,232 ,"TSS24.BF2" ,180 ,1 ,1 ,"打印条形码："
DIRECTION 1,0
BARCODE 16,32 ,"128" ,32 ,1 ,0 ,2 ,2 ,"12345678"
DIRECTION 0,0
PRINT 1, 1
SIZE 40 mm, 30 mm
DIRECTION 0,0
GAP 2.0 mm,0.0 mm
SHIFT 0
DENSITY 7
CLS
TEXT 312,232 ,"TSS24.BF2" ,180 ,1 ,1 ,"打印二维码宽度128："
DIRECTION 1,0
QRCODE 16,32,L,4,A,0,M2,"https://www.xpyun.net"
PRINT 1, 1
	`
	gbk, _ := util.Utf8ToGbk(printContent)
	base64Content := base64.StdEncoding.EncodeToString([]byte(gbk))

	request := model.PrintRequest{}
	request.User = USER_NAME
	request.UserKey = USER_KEY

	//*必填*：打印机编号
	request.Sn = OK_PRINTER_SN

	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = base64Content

	//打印份数，默认为1
	request.Copies = 1

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
	request.Voice = 2

	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1
	result := service.XpYunPos(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}
