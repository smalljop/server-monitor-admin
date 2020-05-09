package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//发送服务器状态信息到服务端
func sendStatToServer() {
	client := &http.Client{}
	stat := GetStat()
	statJson, _ := json.Marshal(stat)
	//url := "http://127.0.0.1:8086/api/client/v1/sync/stat"
	url := "https://smalljop.utools.club/api/client/v1/sync/stat"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(statJson))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	println(string(body))
}
