package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	t_errors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"io"
	"net/http"
	"socketAPI/app/structure"
	"time"
)

func POST(w http.ResponseWriter, r *http.Request, f func(map[string]string) (interface{}, error), log bool) {
	reqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error: %s\n", err)
	}

	querys := r.PostForm
	for key, query := range querys {
		reqData[key] = query[0]
	}

	data, err := f(reqData)
	if log {
		Log("http", reqData, data, err)
	}

	if err != nil {
		res.Code = 1
		if _, ok := err.(*t_errors.TencentCloudSDKError); ok {
			res.Code = 2
		}
		if err.Error() == "event not exist" {
			res.Code = 4
		}
		res.Error = err.Error()
		msg, _ := json.Marshal(res)
		w.Write(msg)
	} else {
		res.Data = data
		msg, _ := json.Marshal(res)

		w.Write(msg)
	}
}

func GET(w http.ResponseWriter, r *http.Request, f func(map[string]string) (interface{}, error), log bool) {
	reqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
	querys := r.URL.Query()
	for key, query := range querys {
		reqData[key] = query[0]
	}
	data, err := f(reqData)
	if log {
		Log("http", reqData, data, err)
	}

	if err != nil {
		w.WriteHeader(400)
		res.Code = 1
		res.Error = err.Error()
		if err.Error() == "machine_code can not be empty" {
			res.Code = 3
		}
		msg, _ := json.Marshal(res)
		w.Write(msg)
	} else {
		res.Data = data
		msg, _ := json.Marshal(res)
		w.Write(msg)
	}
}

func HttpGet(url string, reqParams map[string]string) ([]byte, error) {
	var tr *http.Transport
	tr = &http.Transport{
		MaxIdleConns: 200,
	}

	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)

	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, body)

	params := req.URL.Query()

	for k, reqParam := range reqParams {
		params.Add(k, reqParam)
	}
	req.URL.RawQuery = params.Encode()

	res, err := client.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
