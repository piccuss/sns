package internal

import (
	"sns/internal/pkg"
	"time"
)

func LauchSns(configPath string) {
	config := pkg.LoadConfig(configPath)

	createAlarmCronJob(config)

	pkg.Sugar().Infof("created alarm cron job, just wait for time up!")

	for {
		time.Sleep(time.Second)
	}
}
