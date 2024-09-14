package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xpyun-opensdk/model"
)

/**
 * 发送http的json请求
 *
 * @param url 请求url
 * @param jsonStr 发送的json字符串
 *
 */
func HttpPostJson(url string, data interface{}) *model.XPYunResp {
	b, err := json.Marshal(&data)
	if err != nil {
		var msg = fmt.Sprintf("json serialize err:%+v", err)
		fmt.Println(msg)
		result := model.XPYunResp{
			HttpStatusCode: 500,
		}
		return &result
	}

	var reqMsg = fmt.Sprintf("response result:%+v", string(b))
	fmt.Println(reqMsg)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		var msg = fmt.Sprintf("post json error:%+v", err)
		fmt.Println(msg)
	}

	result := model.XPYunResp{
		HttpStatusCode: resp.StatusCode,
	}

	var content model.XPYunRespContent
	err = json.Unmarshal(body, &content)
	if err == nil {
		result.Content = &content
	} else {
		var msg = fmt.Sprintf("unmarshal body failed, error:%+v", err)
		fmt.Println(msg)
	}

	return &result
}
