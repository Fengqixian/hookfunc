package formatter

import (
	"fmt"
	"strconv"
	"xpyun-opensdk/util"
)

const ROW_MAX_CHAR_LEN = 32
const MAX_NAME_CHAR_LEN = 20
const LAST_ROW_MAX_NAME_CHAR_LEN = 16
const MAX_QUANTITY_CHAR_LEN = 6
const MAX_PRICE_CHAR_LEN = 6

const LAST_ROW_MAX_NAME_CHAR_LEN80 = 24 //26
const ROW_MAX_CHAR_LEN80 = 48
const MAX_NAME_CHAR_LEN80 = 27
const MAX_QUANTITY_CHAR_LEN80 = 7

var orderNameEmpty = util.StrRepeat(" ", MAX_NAME_CHAR_LEN)

// var orderNameEmpty80 = util.StrRepeat(" ", MAX_NAME_CHAR_LEN80)

/**
 * 格式化菜品列表（用于58mm打印机）
 * 注意：默认字体排版，若是字体宽度倍大后不适用
 * 58mm打印机一行可打印32个字符 汉子按照2个字符算
 * 分3列： 名称20字符一般用16字符4空格填充  数量6字符  单价6字符，不足用英文空格填充 名称过长换行
 *
 * @param foodName 菜品名称
 * @param quantity 数量
 * @param price 价格
 * @throws Exception
 */

func FormatPrintOrderItem(foodName string, quantity int, price float64) string {

	foodNameLen := util.CalcGbkLenForPrint(foodName)

	quantityStr := strconv.Itoa(quantity)
	quantityLen := util.CalcAsciiLenForPrint(quantityStr)

	priceStr := fmt.Sprintf("%.2f", price)
	priceLen := util.CalcAsciiLenForPrint(priceStr)

	result := foodName
	mod := foodNameLen % ROW_MAX_CHAR_LEN
	if mod <= LAST_ROW_MAX_NAME_CHAR_LEN {
		// 保证各个列的宽度固定，不足部分，利用空格填充
		result = result + util.StrRepeat(" ", MAX_NAME_CHAR_LEN-mod)

	} else {
		// 另起新行
		result = result + "<BR>"
		result = result + orderNameEmpty
	}

	result = result + quantityStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN-quantityLen)
	result = result + priceStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN-priceLen)

	result = result + "<BR>"

	return result
}

/**
 * 格式化菜品列表（用于80mm打印机）
 * 注意：默认字体排版，若是字体宽度倍大后不适用
 * 80mm打印机一行可打印48个字符 汉子按照2个字符算
 * 分4列： 名称30字符一般用26字符4空格填充  数量6字符  单价6字符 总价6字符，不足用英文空格填充 名称过长换行
 *
 * @param foodName 菜品名称
 * @param quantity 数量
 * @param price 价格
 * @throws Exception
 */

func FormatPrintOrderItem80(foodName string, quantity int, price float64, total float64) string {

	foodNameLen := util.CalcGbkLenForPrint(foodName)

	quantityStr := strconv.Itoa(quantity)
	quantityLen := util.CalcAsciiLenForPrint(quantityStr)

	priceStr := fmt.Sprintf("%.2f", price)
	priceLen := util.CalcAsciiLenForPrint(priceStr)

	totalStr := fmt.Sprintf("%.2f", total)
	totalLen := util.CalcAsciiLenForPrint(totalStr)

	result := foodName
	mod := foodNameLen % ROW_MAX_CHAR_LEN80
	// fmt.Printf("foodNameLen: %d", foodNameLen)
	// fmt.Printf("  mod: %d\r\n", mod)
	if mod <= LAST_ROW_MAX_NAME_CHAR_LEN80 {
		// 保证各个列的宽度固定，不足部分，利用空格填充
		result = result + util.StrRepeat(" ", MAX_NAME_CHAR_LEN80-mod)
		result = result + quantityStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-quantityLen)
		result = result + priceStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-priceLen)
		result = result + totalStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-totalLen)
	} else {
		// 另起新行
		// 字符串分割 一行打印12个汉字，换行
		foods := splitStrArray(foodName, 12)
		tempStr := ""
		for i := 0; i < len(foods); i++ {
			if i == 0 {
				mod := util.CalcGbkLenForPrint(foods[i])
				tempStr = tempStr + foods[i] + util.StrRepeat(" ", MAX_NAME_CHAR_LEN80-mod)
				tempStr = tempStr + quantityStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-quantityLen)
				tempStr = tempStr + priceStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-priceLen)
				tempStr = tempStr + totalStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-totalLen)
				tempStr = tempStr + "<BR>"
			} else if i == len(foods)-1 {
				tempStr = tempStr + foods[i]
			} else {
				tempStr = tempStr + foods[i] + "<BR>"
			}
		}
		result = tempStr
	}

	// result = result + quantityStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-quantityLen)
	// result = result + priceStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-priceLen)
	// result = result + totalStr + util.StrRepeat(" ", MAX_QUANTITY_CHAR_LEN80-totalLen)

	result = result + "<BR>"

	return result
}

func splitStrArray(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}
