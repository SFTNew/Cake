package config

import (
	"encoding/json"
	"os"
)

//服务端配置
type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `josn:"static_path"`
	Mode       string `json:"mode"`
}

var SerConfig AppConfig

func InitConfig() *AppConfig {
	file, err := os.Open("D:/goProject/cake/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
