package main

import "xpyun-opensdk/chinese"

func main() {

	// 1、添加打印机到开发者账户（可批量） 【必接】
	// chinese.AddPrintersTest()

	// 2、设置打印机语音类型
	// chinese.SetVoiceTypeTest()

	// ============3、打印小票订单 begin==========
	// chinese.PrintFontAlign()
	// chinese.PrintFontAlignVoiceSupport()

	// 58小票综合排版样例
	// 不支持金额播报
	// chinese.PrintComplexReceipt()
	// 支持金额播报
	// chinese.PrintComplexReceiptVoiceSupport()

	// 80小票综合排版样例
	// 不支持金额播报
	// chinese.PrintComplexReceipt80()
	// 支持金额播报
	// chinese.PrintComplexReceiptVoiceSupport80()
	// ============3、打印小票订单 end==========

	//  ============4、打印标签订单 begin============
	// 标签综合排版样例
	// chinese.PrintLabel()
	// ============4、打印标签订单 end==========

	// 5、批量删除打印机
	// chinese.DelPrintersTest()

	// 6、修改打印机信息
	// chinese.UpdPrinterTest()

	// 7、清空待打印队列
	// chinese.XpYunDelPrinterQueueTest()

	// ============数据查询类接口 begin=========
	// 8、查询订单是否打印成功
	// chinese.XpYunQueryOrderStateTest()

	// 9、查询指定打印机某天的订单统计数
	// chinese.XpYunQueryOrderStatisTest()

	// 10、获取指定打印机状态
	chinese.XpYunQueryPrinterStatusTest()

	// 11、批量获取指定打印机状态
	// chinese.XpYunQueryPrintersStatusTest()
	// ============数据查询类接口 end=========

	// 12、金额播报
	// chinese.XpYunPlayVoiceTest()

	// =======NEW=======
	// // 13、POS 58 小票模板指令打印
	// chinese.XpYunPosTest()
	// // 13、POS 80 小票模板指令打印
	// chinese.XpYunPosTest80()
	// // 13、POS 58 标签模板指令打印
	// chinese.XpYunPosTsplTest()

	// // 14、钱箱控制
	// chinese.XpYunControlBoxTest()

	// // 15、扩展语音播报
	// chinese.XpYunPlayVoiceExtTest()

	// // 16、自定义语音播报
	// chinese.XpYunPlayCustomVoiceTest()

	// // 17、店铺 LOGO 上传
	// chinese.XpYunUploadLogoTest()

	// // 18、店铺 LOGO 删除
	// chinese.XpYunDelUploadLogoTest()

	// // 19、获取打印机基本信息
	// chinese.XpYunPrinterInfoTest()
}
