package main

// 配置
type Config struct {
	Debug     bool   `json:"debug"`     //Debug模式
	ServerUrl string `json:"serverUrl"` //服务端地址包含端口
	Version   string `json:"version"`   //当前版本

}
