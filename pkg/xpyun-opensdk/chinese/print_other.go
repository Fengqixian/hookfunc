package chinese

//票据标签使用示例
import (
	"encoding/json"
	"fmt"
	"xpyun-opensdk/model"
	"xpyun-opensdk/service"
)

/**
 * 小票打印字体对齐样例，不支持金额播报
 * 注意：对齐标签L C R CB 请勿嵌套使用，嵌套使用内层标签有效，外层失效；
 * 同一行请勿使用多个对齐标签，否则只有最后一个对齐标签有效
 */
func PrintFontAlign() {

	/**
	 <BR>：换行符（同一行有闭合标签(如 </C> 或 </R> )则应放到闭合标签前面, 连续两个换行符<BR><BR>可以表示加一空行）
	 <L></L>：左对齐
	 <C></C>：居中对齐
	 <R></R>：右对齐
	 注意：居中（或居右）标签生效需要把 <BR> 标签放在 </CB> 或 </C> 或 </R> 标签前面。同一行内容不能使用多种对齐方式，
	可通过补空格方式自定义对齐样式。与换行标签一起使用时，换行标签应放在对齐标签里面。
	使用举例说明：
	      右对齐且换行  <R>待右对齐内容<BR></R>
	      58mm的机器，一行打印16个汉字，32个字母
	      80mm的机器，一行打印24个汉字，48个字母

	 <N></N>：字体正常大小
	 <HB></HB>：字体变高一倍
	 <WB></WB>：字体变宽一倍
	 <B></B>：字体放大一倍
	 <CB></CB>：字体放大一倍居中
	 <CUT>：自定义切刀
	 <HB2></HB2>：字体变高二倍
	 <WB2></WB2>：字体变宽二倍
	 <B2></B2>：字体放大二倍
	 <BOLD></BOLD>：字体加粗
	 <IMG></IMG>：打印LOGO图片，需登录开放平台在【打印管理➜设备管理】下通过设置LOGO功能进行上传。此处直接写入空标
	签,如 <IMG></IMG> 即可, 具体可参考样例。图片宽度设置：可以通过 <IMG> 标签名称自定义，如 <IMG60> 表示宽度为60，
	相应的闭合标签 </IMG>不需要指定高度。<IMG> 标签不指定宽度默认为40，最小值为20，最大值为200，图片数据最大值30KB
	 <QRCODE s=6 e=L l=center>二维码内容</QRCODE>：二维码（标签内容是二维码值, 最大不能超过256个字符）。
	      s：二维码大小，默认6，取值范围 0-16，当s为0时，二维码大小会根据内容大小动态调整
	      e：二维码纠错等级，默认L，取值范围 L/M/Q/H
	      l：二维码打印位置，默认center，取值范围 left/center/right。（l为location位置单词首字母）
	 <FONT w="0" h="0">需要放大的文档内容</FONT>：字体标签，字体倍数放大。
	      属性w: 字体宽度放大倍数。默认0，不放大，取值[0-7]
	      属性h: 字体高度放大倍数。默认0，不放大，取值[0-7]
	 <BARCODE t=CODE128 w=2 h=100 p=2>条形码值</BARCODE>：条形码（标签内容是条形码值）。若不能正常打印一维码，则需
	在<BARCODE>前面加上<BR>标签
	      t：条码类型，取值范围：UPCA、EAN13、EAN8、CODE39、ITF、CODABAR、CODE93、CODE128。默认值：CODE128
	      w：条码宽度，取值范围：2-6；默认值：2
	      h：条码高度，取值范围：1-255；默认值：100
	      p：条码值显示的位置，取值范围：0不显示，1条码上方，2条码下方，3条码上下方都显示；默认值：2
	 <RH n="3">放置需更改行间距内容</RH>：小票内容行间距设置标签，只有内容达到两行或两行以上才能生效。<RH></RH>之间的
	内容是要调整行间距的小票内容，内容可以是小票其他标签包裹的内容，也可以是纯文本内容。RH标签需要放在基础标签<L></L>、
	<C></C>或<R></R>之间才能生效。若有使用<BOLD>标签时,<RH>标签需放在<BOLD>标签里面。放置方式可参考下发的使用举例说
	明。参数说明如下：
	    n：行间距值，n的取值范围[0,5]，设置时最好让n乘以48是一个非负整数，若乘积出现小数位时，将舍弃小数位取
	整数。默认行间距高是48即约高3.75mm
	  使用举例说明：
	    打印行间距为3且内容加粗的标签使用示例
	<R><BOLD><RH n="3">欢迎使用芯烨云打印机，RH标签是小票内容行间距设置标签。需要放在基础标签L、C或R之间才能生效哟！
	</RH></BOLD></R><BR>
	 <P></P>：开启横向左右布局模式。需结合AREA标签使用。仅用于支持左右布局的机器，详情可加开放平台QQ技术群或企业微信
	群垂询
	 <AREA></AREA>：按区域打印，需结合P标签使用。仅用于支持左右布局的机器，详情可加开放平台QQ技术群或企业微信群垂询。
	      x:打印区域起始点的x轴；
	      y:打印区域起始点的y轴；
	      w:打印区域的宽度；
	      h:打印区域的高度。
	  使用举例说明：<P><AREA x=0 y=0 w=100 h=200>该区域需要打印的内容，区域内部文字会自动换行</AREA></P>
	*/

	printContent := `不加标签：默认字体大小<BR>
<BR>
L标签：<L>左对齐<BR></L>
<BR>
R标签：<R>右对齐<BR></R>
<BR>
C标签：<C>居中对齐<BR></C>
<BR>
N标签：<N>字体正常大小<BR></N>
<BR>
HB标签：<HB>字体变高一倍<BR></HB>
<BR>
WB标签：<WB>字体变宽一倍<BR></WB>
<BR>
B标签：<B>字体放大一倍<BR></B>
<BR>
HB2标签：<HB2>字体变高二倍<BR></HB2>
<BR>
WB2标签：<WB2>字体变宽二倍<BR></WB2>
<BR>
B2标签：<B2>字体放大二倍<BR></B2>
<BR>
BOLD标签：<BOLD>字体加粗<BR></BOLD>`

	printContent = printContent + "<BR>"
	// 嵌套使用对齐和字体
	printContent = printContent + "<C>嵌套使用：<BOLD>居中加粗</BOLD><BR></C>"

	// 打印条形码和二维码
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><BARCODE t=CODE128 w=2 h=100 p=2>9884822189</BARCODE></C>"
	printContent = printContent + "<C><QRCODE s=6 e=L l=center>https://www.xpyun.net</QRCODE></C>"

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

	result := service.XpYunPrint(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}

/**
 * 小票打印字体对齐样例，支持金额播报
 * 注意：对齐标签L C R CB 请勿嵌套使用，嵌套使用内层标签有效，外层失效；
 * 同一行请勿使用多个对齐标签，否则只有最后一个对齐标签有效
 */
func PrintFontAlignVoiceSupport() {

	/**
	 <BR>：换行符（同一行有闭合标签(如 </C> 或 </R> )则应放到闭合标签前面, 连续两个换行符<BR><BR>可以表示加一空行）
	 <L></L>：左对齐
	 <C></C>：居中对齐
	 <R></R>：右对齐
	 注意：居中（或居右）标签生效需要把 <BR> 标签放在 </CB> 或 </C> 或 </R> 标签前面。同一行内容不能使用多种对齐方式，
	可通过补空格方式自定义对齐样式。与换行标签一起使用时，换行标签应放在对齐标签里面。
	使用举例说明：
	      右对齐且换行  <R>待右对齐内容<BR></R>
	      58mm的机器，一行打印16个汉字，32个字母
	      80mm的机器，一行打印24个汉字，48个字母

	 <N></N>：字体正常大小
	 <HB></HB>：字体变高一倍
	 <WB></WB>：字体变宽一倍
	 <B></B>：字体放大一倍
	 <CB></CB>：字体放大一倍居中
	 <CUT>：自定义切刀
	 <HB2></HB2>：字体变高二倍
	 <WB2></WB2>：字体变宽二倍
	 <B2></B2>：字体放大二倍
	 <BOLD></BOLD>：字体加粗
	 <IMG></IMG>：打印LOGO图片，需登录开放平台在【打印管理➜设备管理】下通过设置LOGO功能进行上传。此处直接写入空标
	签,如 <IMG></IMG> 即可, 具体可参考样例。图片宽度设置：可以通过 <IMG> 标签名称自定义，如 <IMG60> 表示宽度为60，
	相应的闭合标签 </IMG>不需要指定高度。<IMG> 标签不指定宽度默认为40，最小值为20，最大值为200，图片数据最大值30KB
	 <QRCODE s=6 e=L l=center>二维码内容</QRCODE>：二维码（标签内容是二维码值, 最大不能超过256个字符）。
	      s：二维码大小，默认6，取值范围 0-16，当s为0时，二维码大小会根据内容大小动态调整
	      e：二维码纠错等级，默认L，取值范围 L/M/Q/H
	      l：二维码打印位置，默认center，取值范围 left/center/right。（l为location位置单词首字母）
	 <FONT w="0" h="0">需要放大的文档内容</FONT>：字体标签，字体倍数放大。
	      属性w: 字体宽度放大倍数。默认0，不放大，取值[0-7]
	      属性h: 字体高度放大倍数。默认0，不放大，取值[0-7]
	 <BARCODE t=CODE128 w=2 h=100 p=2>条形码值</BARCODE>：条形码（标签内容是条形码值）。若不能正常打印一维码，则需
	在<BARCODE>前面加上<BR>标签
	      t：条码类型，取值范围：UPCA、EAN13、EAN8、CODE39、ITF、CODABAR、CODE93、CODE128。默认值：CODE128
	      w：条码宽度，取值范围：2-6；默认值：2
	      h：条码高度，取值范围：1-255；默认值：100
	      p：条码值显示的位置，取值范围：0不显示，1条码上方，2条码下方，3条码上下方都显示；默认值：2
	 <RH n="3">放置需更改行间距内容</RH>：小票内容行间距设置标签，只有内容达到两行或两行以上才能生效。<RH></RH>之间的
	内容是要调整行间距的小票内容，内容可以是小票其他标签包裹的内容，也可以是纯文本内容。RH标签需要放在基础标签<L></L>、
	<C></C>或<R></R>之间才能生效。若有使用<BOLD>标签时,<RH>标签需放在<BOLD>标签里面。放置方式可参考下发的使用举例说
	明。参数说明如下：
	    n：行间距值，n的取值范围[0,5]，设置时最好让n乘以48是一个非负整数，若乘积出现小数位时，将舍弃小数位取
	整数。默认行间距高是48即约高3.75mm
	  使用举例说明：
	    打印行间距为3且内容加粗的标签使用示例
	<R><BOLD><RH n="3">欢迎使用芯烨云打印机，RH标签是小票内容行间距设置标签。需要放在基础标签L、C或R之间才能生效哟！
	</RH></BOLD></R><BR>
	 <P></P>：开启横向左右布局模式。需结合AREA标签使用。仅用于支持左右布局的机器，详情可加开放平台QQ技术群或企业微信
	群垂询
	 <AREA></AREA>：按区域打印，需结合P标签使用。仅用于支持左右布局的机器，详情可加开放平台QQ技术群或企业微信群垂询。
	      x:打印区域起始点的x轴；
	      y:打印区域起始点的y轴；
	      w:打印区域的宽度；
	      h:打印区域的高度。
	  使用举例说明：<P><AREA x=0 y=0 w=100 h=200>该区域需要打印的内容，区域内部文字会自动换行</AREA></P>
	*/

	printContent := `不加标签：默认字体大小<BR>
<BR>
L标签：<L>左对齐<BR></L>
<BR>
R标签：<R>右对齐<BR></R>
<BR>
C标签：<C>居中对齐<BR></C>
<BR>
N标签：<N>字体正常大小<BR></N>
<BR>
HB标签：<HB>字体变高一倍<BR></HB>
<BR>
WB标签：<WB>字体变宽一倍<BR></WB>
<BR>
B标签：<B>字体放大一倍<BR></B>
<BR>
HB2标签：<HB2>字体变高二倍<BR></HB2>
<BR>
WB2标签：<WB2>字体变宽二倍<BR></WB2>
<BR>
B2标签：<B2>字体放大二倍<BR></B2>
<BR>
BOLD标签：<BOLD>字体加粗<BR></BOLD>`

	printContent = printContent + "<BR>"
	// 嵌套使用对齐和字体
	printContent = printContent + "<C>嵌套使用：<BOLD>居中加粗</BOLD><BR></C>"

	// 打印条形码和二维码
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><BARCODE t=CODE128 w=2 h=100 p=2>9884822189</BARCODE></C>"
	printContent = printContent + "<C><QRCODE s=6 e=L l=center>https://www.xpyun.net</QRCODE></C>"

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
	request.Voice = 1

	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1

	//支付方式：
	//取值范围41~55：
	//支付宝 41、微信 42、云支付 43、银联刷卡 44、银联支付 45、会员卡消费 46、会员卡充值 47、翼支付 48、成功收款 49、嘉联支付 50、壹钱包 51、京东支付 52、快钱支付 53、威支付 54、享钱支付 55
	//仅用于支持金额播报的芯烨云打印机。
	request.PayType = 41

	//支付与否：
	//取值范围59~61：
	//退款 59 到账 60 消费 61。
	//仅用于支持金额播报的芯烨云打印机。
	request.PayMode = 60

	//支付金额：
	//最多允许保留2位小数。
	//仅用于支持金额播报的芯烨云打印机。
	request.Money = 20.15

	result := service.XpYunPrint(&request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))
	fmt.Println(msg)
}
