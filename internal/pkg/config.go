package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	StockCodes      []string `json:"stock_codes"`
	AlarmTime       []string `json:"alarm_time"`
	StockDataSource string   `json:"stock_data_srouce"`
	Reciever        string   `json:"reciever"`
	NotifyVendor    string   `json:"notify_vendor"`
}

func (s Config) String() string {
	return fmt.Sprintf("{stockCodes=%s, alarmTime=%s, stockDataSource=%s, reciever=%s, notifyVendor=%s}",
		s.StockCodes, s.AlarmTime, s.StockDataSource, s.Reciever, s.NotifyVendor)
}

func LoadConfig(configPath string) Config {
	configStr, err := ioutil.ReadFile(configPath)
	if err != nil {
		Sugar().Fatalf("read config file failed, configPath=%s, ", configPath, err)
	}
	var config Config
	err = json.Unmarshal(configStr, &config)
	if err != nil {
		Sugar().Fatalf("config parse json fialed. config=%s", configStr)
	}
	if ok := checkConfigValid(config); !ok {
		Sugar().Fatalf("config is invalid. config=%s", config)
	}
	Sugar().Infof("load config success. config=%s", config)
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
