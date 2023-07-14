package internal

import (
	"sns/internal/datasource"
	"sns/internal/log"
	"time"
)

func LauchSns(configPath string) {
	config := loadConfig(configPath)
	startService(config)
}

func startService(config Config) {
	datasource := datasource.GetStockDataSource(config.StockDataSource)
	if datasource == nil {
		log.Sugar().Fatalf("getStockDataSource is nil. name=%s", config.StockDataSource)
	}
	log.Sugar().Infof("use StockDataSource: %s", config.StockDataSource)

	createCronJob(config, datasource)
	log.Sugar().Infof("created cron job, just wait for time up!")

	for {
		time.Sleep(time.Second)
	}
}
