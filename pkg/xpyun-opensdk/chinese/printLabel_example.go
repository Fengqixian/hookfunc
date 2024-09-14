package chinese

//标签综合排版打印样例
import (
	"encoding/json"
	"fmt"
	"xpyun-opensdk/model"
	"xpyun-opensdk/service"
	"xpyun-opensdk/util"
)

/**
 * 标签打印综合排版样例
 * 如何确定坐标：坐标原点位于左上角，x轴是从左往右，y轴是从上往下；
 * 根据测试，x轴最大值=标签纸宽度*8，y轴最大值=标签纸高度*8，
 * 如标签纸尺寸为40*30，x轴最大值=40*8=320，y轴最大值=30*8=240
 * 实际排版效果需要用户按实际纸张尺寸和需求自行排版
 *
 * 打印内容内（标签除外）大于号和小于号需要经过转译才能正常打印。其中，“<”用“&lt”表示，“>”用“&gt”表示；1mm=8dots。
 */
func PrintLabel() {
	/**
	<PAGE n="2" l="2"></PAGE>：
		分页，用于支持打印多张不同的标签页面，不使用该标签表示所有元素只打印在一个标签页
		属性 n 为当前页打印份数，必须为正整数，当n不指定时，默认为打印1份
		属性 l 为当前两个标签纸之间的间隙。当l不指定时，默认为2mm

	<SIZE>width,height</SIZE>：
		设置标签纸宽高，width 标签纸宽度(不含背纸)，height 标签纸高度(不含背纸)，单位mm，如<SIZE>40,30</SIZE> ，且
		必须放到PAGE标签页里面才生效
		以下属性x、y均以dot为单位。例如标签纸大小为：宽40mm，高度为30mm。1mm=8dot
	    x的最大值就是40*8=320，x的取值范围是0-320。
	    y的最大值就是为30*8=240，y的取值范围是0-240。

	<TEXT x="10" y="100" font="9" w="1" h="2" r="0">文本内容</TEXT>：
		打印文本，其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 font 为字体，仅支持以下参数:
	     1、 8 x 12 dot 英数字体
	     2、 12 x 20 dot英数字体
	     3、 16 x 24 dot英数字体
	     4、 24 x 32 dot英数字体
	     5、 32 x 48 dot英数字体
	     6、 14 x 19 dot英数字体 OCR-B
	     7、 21 x 27 dot 英数字体OCR-B
	     8、 14 x25 dot英数字体OCR-A
	     9、 简体中文 24dot x 24dot font (GB 码)(3mm X 3mm)
		属性 w 为文字宽度放大倍率1-10（默认为1）
		属性 h 为文字高度放大倍率1-10（默认为1）
		属性 r 为文字旋转角度(顺时针方向，默认为0)，仅支持以下参数值设置：
	     0     0度
	     90   90度
	     180 180度
	     270度

	<BC128 x="10" y="100" h="60" s="1" n="1" w="1" r="0">1234567</BC128>：
		打印code128一维码，其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 h 为条形码高度（默认为48）
		属性 s 是否人眼可识：0 不可识，1 可识（默认为1）
		属性 n 为窄 bar 宽度，以点(dot)表示（默认为1）
		属性 w 为宽 bar 宽度，以点(dot)表示（默认为1）
		属性 r 为文字旋转角度 (顺时针方向，默认为0)，仅支持以下参数值设置：
	     0     0度
	     90   90度
	     180 180度
	     270 270度

	<BC39 x="10" y="100" h="60" s="1" n="1" w="1" r="0">1234567</BC39>：
		打印code39一维码，其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 h 为条形码高度（默认为48）
		属性 s 是否人眼可识：0 不可识，1 可识（默认为1）
		属性 n 为窄 bar 宽度，以点(dot)表示（默认为1）
		属性 w 为宽 bar 宽度，以点(dot)表示（默认为2）
		属性 r 为文字旋转角度(顺时针方向，默认为0)，仅支持以下参数值设置：
	     0     0度
	     90   90度
	     180 180度
	     270度

	<QRC x="20" y="20" s="2" e="L">二维码内容</QRC>：
		打印二维码，标签内容是二维码值, 最大不能超过256个字符，其中：
		属性 x 为水平方向起始点坐标(默认为0)，取值范围（0至标签宽度*8）1 mm=8 dots
		属性 y 为垂直方向起始点坐标(默认为0)，取值范围（0至标签高度*8）1 mm=8 dots
		属性 s 为二维码打印大小(默认为2 取值1-10)
		属性 e 为二维码纠错等级(默认为L 取值L/M/Q/H)

	<IMG x="16" y="32" w="100">：
		打印LOGO图片，需登录开放平台在【打印管理➜设备管理】下通过设置LOGO功能进行上传。此处直接
		写入空标签,若指定了<PAGE>标签，<IMG>标签应该放到<PAGE>标签里面， <IMG>, 如 <IMG>即可, 具
		体可参考样例。其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 w 为logo图片最大宽度（默认为50），最小值为20，最大值为100。logo图片的高度和宽度相等

	<SEQ x="8" y="8" xe="100" ye="100" s="4">：
		打印方框。其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 xe 为水平方向结束点坐标（默认为0）
		属性 ye 为垂直方向结束点坐标（默认为0）
		属性 s 为方框线条打印大小（默认为4 取值1-10）。

	<L x="20" y="20" w="4" h="250">：
		打印表格线。其中：
		属性 x 为水平方向起始点坐标（默认为0）
		属性 y 为垂直方向起始点坐标（默认为0）
		属性 w 为线条的宽度（默认为4 取值1-1000）
		属性 h 为线条的高度（默认为4 取值1-1000）
	*/

	//第一个标签
	printContent := "<PAGE>"
	// 设定标签纸尺寸
	printContent = printContent + "<SIZE>40,30</SIZE>"
	printContent = printContent + `<TEXT x="8" y="8" w="1" h="1" r="0">` + `#001` + util.StrRepeat(" ", 4) + "一号桌" + util.StrRepeat(" ", 4) + "1/3" + "</TEXT>" + `<TEXT x="8" y="96" w="2" h="2" r="0">` + `黄金炒饭` + `</TEXT>` + `<TEXT x="8" y="200" w="1" h="1" r="0">` + `王女士` + util.StrRepeat(" ", 4) + `136****3388` + `</TEXT>` + `</PAGE>`

	//第二个标签
	printContent = printContent + "<PAGE>"
	printContent = printContent + `<TEXT x="8" y="8" w="1" h="1" r="0">` + "#001" + util.StrRepeat(" ", 4) +
		`一号桌` + util.StrRepeat(" ", 4) +
		`2/3` +
		`</TEXT>` +
		`<TEXT x="8" y="96" w="2" h="2" r="0">` +
		`凉拌青瓜` +
		`</TEXT>` +
		`<TEXT x="8" y="200" w="1" h="1" r="0">` +
		`王女士` + util.StrRepeat(" ", 4) +
		`136****3388` +
		"</TEXT>" +
		"</PAGE>"

	//第三个标签
	printContent = printContent + "<PAGE>"
	printContent = printContent + `<TEXT x="8" y="8" w="1" h="1" r="0">` +
		"#001" + util.StrRepeat(" ", 4) +
		"一号桌" + util.StrRepeat(" ", 4) +
		"3/3" +
		"</TEXT>" +
		`<TEXT x="8" y="96" w="2" h="2" r="0">` +
		"老刘家肉夹馍" +
		"</TEXT>" +
		`<TEXT x="8" y="200" w="1" h="1" r="0">` +
		"王女士" + util.StrRepeat(" ", 4) +
		"136****3388" +
		"</TEXT>" +
		"</PAGE>"

	//第四个标签 打印条形码
	printContent = printContent + "<PAGE>"
	printContent = printContent + `<TEXT x="8" y="8" w="1" h="1" r="0">` +
		"打印条形码：" +
		"</TEXT>" +
		`<BC128 x="16" y="32" h="32" s="1" n="2" w="2" r="0">` +
		"12345678" +
		"</BC128>" +
		"</PAGE>"

	//第四个标签 打印二维码，宽度最小为128 低于128会无法扫描
	printContent = printContent + "<PAGE>"
	printContent = printContent + `<TEXT x="8" y="8" w="1" h="1" r="0">` +
		"打印二维码宽度128：" + "</TEXT>" +
		`<QRC x="16" y="32" s="2" e="L">` +
		"https://www.xpyun.net" +
		"</QRC>" +
		"</PAGE>"
		//第五个标签 打印表格
	printContent = printContent + "<PAGE><SIZE>40,30</SIZE>"
	printContent = printContent + `<SEQ x="8" y="8" xe="296" ye="232" s="3">` +
		`<L x="88" y="8" w="3" h="224">` +
		`<TEXT x="24" y="16" w="1" h="1" r="0">编号</TEXT>` +
		`<TEXT x="96" y="16" w="1" h="1" r="0">20220419</TEXT>` +
		`<L x="8" y="48" w="284" h="3">` +
		`<TEXT x="24" y="56" w="1" h="1" r="0">名称</TEXT>` +
		`<TEXT x="96" y="56" w="1" h="1" r="0">芯烨云</TEXT>` +
		`<L x="8" y="88" w="284" h="3">` +
		`<TEXT x="24" y="96" w="1" h="1" r="0">规格</TEXT>` +
		`<TEXT x="96" y="96" w="1" h="1" r="0">V3.90</TEXT>` +
		`<L x="8" y="128" w="284" h="3">` +
		`<TEXT x="24" y="136" w="1" h="1" r="0">数量</TEXT>` +
		`<TEXT x="96" y="136" w="1" h="1" r="0">1</TEXT>` +
		`<L x="8" y="168" w="180" h="3">` +
		`<TEXT x="24" y="176" w="1" h="1" r="0">日期</TEXT>` +
		`<TEXT x="96" y="176" w="1" h="1" r="0">04/19</TEXT>` +
		`<L x="184" y="128" w="3" h="104">` +
		`<QRC x="195" y="136" s="3" e="L">https://www.xpyun.net/open/index.html</QRC>` +
		"</PAGE>"

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
	result := service.XpYunPrintLabel(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}
