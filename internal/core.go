package internal

import (
	"sns/internal/datasource"
	"sns/internal/notify"
	"sns/internal/pkg"

	"github.com/robfig/cron"
)

func createAlarmCronJob(config pkg.Config) {
	alarmTimes := config.AlarmTime
	c := cron.New()
	for _, time := range alarmTimes {
		c.AddFunc(time, func() {
			invokeAlaramProc(config)
		})
		pkg.Sugar().Infof("alarmCronJob created, time=%s", time)
	}
}

func invokeAlaramProc(config pkg.Config) {
	pkg.Sugar().Infof("invokeAlaramProc...")
	stockData := datasource.FetchStockData(config)
	err := notify.SendEmail(config, stockData)
	if err != nil {
		pkg.Sugar().Errorf("sendEmail failed. err=%s", &err)
	}
	pkg.Sugar().Infof("endAlaramProc...")
}
