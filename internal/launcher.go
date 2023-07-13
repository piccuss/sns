package internal

import (
	_ "sns/internal/datasource"
)

func LauchSns(configPath string) {
	config := loadConfig(configPath)
	startService(config)
}

func startService(config Config) {
	//获取数据源与解析器
	//启动cron任务
	//等待
}
