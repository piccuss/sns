package internal

import (
	"sns/internal/datasource"
	"sns/internal/pkg"

	"github.com/robfig/cron"
)

func createAlarmCronJob(config pkg.Config, datasource datasource.StockDataSource) {
	//TODO start multi cron job
	alarmTimes := config.AlarmTime
	c := cron.New()
	for _, time := range alarmTimes {
		c.AddFunc(time, func() {

		})

	}
}
