package internal

import (
	"encoding/json"
	"io/ioutil"
	"sns/internal/log"
)

type Config struct {
	StockCodes []string `json:"stock_codes"`
	AlarmTime  []string `json:"alarm_time"`
}

func loadConfig(configPath string) Config {
	configStr, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Sugar().Fatalf("read config file failed, configPath=%s, ", configPath, err)
	}
	var config Config
	err = json.Unmarshal(configStr, &config)
	if err != nil {
		log.Sugar().Fatalf("config parse json fialed. config=%s", configStr)
	}
	log.Sugar().Infof("load config success. config=%s", config)
	return config
}
