package internal

import (
	"sns/internal/datasource"
	"sns/internal/pkg"
	"time"
)

func LauchSns(configPath string) {
	config := pkg.LoadConfig(configPath)
	startService(config)
}

func startService(config pkg.Config) {
	datasource := datasource.GetStockDataSource(config.StockDataSource)
	if datasource == nil {
		pkg.Sugar().Fatalf("getStockDataSource is nil. name=%s", config.StockDataSource)
	}
	pkg.Sugar().Infof("use StockDataSource: %s", config.StockDataSource)

	createAlarmCronJob(config, datasource)
	pkg.Sugar().Infof("created alarm cron job, just wait for time up!")

	for {
		time.Sleep(time.Second)
	}
}
