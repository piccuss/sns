package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	StockCodes     []string `json:"stock_codes"`
	AlarmTime      []string `json:"alarm_time"`
	StaggingTipsOn bool     `json:"stagging_tips_on"`
}

var config Config

// LoadConfig from local config json file
func LoadConfig() {
	jsc, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Fatal("read config.json error", err)
	}
	err = json.Unmarshal(jsc, &config)
	if err != nil {
		log.Fatal("unmarshal config.json error", err)
	}
	log.Println("load config:", config)
}
