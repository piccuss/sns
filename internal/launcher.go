package internal

import (
	"sns/internal/datasource"
	_ "sns/internal/datasource"
	"sns/internal/log"
)

func LauchSns(configPath string) {
	config := loadConfig(configPath)
	startService(config)
}

func startService(config Config) {
	//获取数据源与解析器
	datasource := datasource.GetStockDataSource(config.StockDataSource)
	if datasource == nil {
		log.Sugar().Fatalf("getStockDataSource is nil. name=%s", config.StockDataSource)
	}
	//启动cron任务
	//等待
}
