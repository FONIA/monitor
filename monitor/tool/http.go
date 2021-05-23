package tool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) bool {

	// 超时时间：5秒
	client := &http.Client{Timeout:5 * time.Second}
	_, err := client.Get(url)
	if err != nil {
		return false
	}
	return true

}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func GetData(url string)(int,string)  {
	// 超时时间：5秒
	client := &http.Client{Timeout:5 * time.Second}
	con, err := client.Get(url)
	if err != nil {
		return 0,""
	}
	return con.StatusCode,con.Status

}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "err"
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}