package internal

import (
	"encoding/json"
	"io/ioutil"
	"sns/internal/log"
)

type Config struct {
	StockCodes      []string `json:"stock_codes"`
	AlarmTime       []string `json:"alarm_time"`
	StockDataSource string   `json:"stock_data_srouce"`
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
	if ok := checkConfigValid(config); !ok {
		log.Sugar().Fatalf("config is invalid. config=%s", config)
	}
	log.Sugar().Infof("load config success. config=%s", config)
	return config
}

func checkConfigValid(config Config) (ok bool) {
	ok = true
	if len(config.StockCodes) == 0 || len(config.AlarmTime) == 0 || len(config.StockDataSource) == 0 {
		ok = false
		return
	}
	return
}
