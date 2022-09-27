/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-09-27 15:01:39
 * @LastEditors: solid
 * @LastEditTime: 2022-09-27 16:05:26
 */
package utils

import (
	"CSDN/conf"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost: 8,
	}
	client = &http.Client{
		Transport: tr,
	}
)

func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("导出成功:", name)
	}
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//发送请求
func SendMessageServer(remoteURL string, data []byte) []byte {
	var req *http.Request
	var err error
	if data == nil {
		req, err = http.NewRequest("GET", remoteURL, nil)
	} else {
		req, err = http.NewRequest("POST", remoteURL, bytes.NewBuffer(data))
	}
	if err != nil {
		fmt.Println("http.GET error: ", err.Error())
		return nil
	}
	req.Header.Set("cookie", conf.Config.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http.GET error: ", err.Error())
		return nil
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
