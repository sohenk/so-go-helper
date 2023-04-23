package httputils

import (
	"io/ioutil"
	"net/http"
	"time"
)

// 发送请求
func SendRequest(method string, url string, body []byte, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if timeout == 0 {
		timeout = time.Millisecond * 200
	}
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
