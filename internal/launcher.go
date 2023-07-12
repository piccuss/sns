package internal

// 启动sns服务
func LauchSns(configPath string) {
	//加载sns配置
	config := loadConfig(configPath)
	//启动sns任务
	startService(config)
}

func startService(config Config) {
	//获取数据源与解析器
	//启动cron任务
	//等待
}
