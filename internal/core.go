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
	c.Start()
}

func invokeAlaramProc(config pkg.Config) {
	pkg.Sugar().Infof("invokeAlaramProc...")
	stockData, err := datasource.FetchStockData(config.StockDataSource, config.StockCodes)
	if err != nil {
		pkg.Sugar().Errorf("fetchStockData failed. err=%s", &err)
		return
	}
	err = notify.SendEmail(config.Reciever, config.NotifyVendor, stockData)
	if err != nil {
		pkg.Sugar().Errorf("sendEmail failed. err=%s", &err)
		return
	}
	pkg.Sugar().Infof("endAlaramProc...")
}
