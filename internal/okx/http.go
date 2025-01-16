package okx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpOkx struct {
	BaseURL string
	Retry   int
}

func NewHttpOkx(baseURL string, retry int) *HttpOkx {
	return &HttpOkx{BaseURL: baseURL, Retry: retry}
}

func (h *HttpOkx) Get(url string, v any) error {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, h.BaseURL+url, nil)
	if err != nil {
		return err
	}

	count := 0
	var response http.Response
	for {
		res, err := client.Do(req)
		if err != nil {
			count++
			if count >= h.Retry {
				return err
			}
			continue
		}

		response = *res
		break
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func (h *HttpOkx) Post(url string, data interface{}, v any) error {
	// 将数据编码为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, h.BaseURL+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	count := 0
	var response http.Response
	for {
		res, err := client.Do(req)
		if err != nil {
			count++
			if count >= h.Retry {
				return err
			}

			continue
		}

		response = *res
		break
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
