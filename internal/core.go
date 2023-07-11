package internal

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	queryUrlPrefix = "https://hq.sinajs.cn/list="
)

var ctsZone = time.FixedZone("UTC+8", 8*60*60)

func Start() {
	//1.加载配置到内存
	LoadConfig()
	//2.按配置启动定时流程
	for i := 0; i < len(config.AlarmTime); i++ {
		go launchJob(config.AlarmTime[i], config.StockCodes)
	}
	select {}
}

func launchJob(alarmTime string, codes []string) {
	jobId := "J" + strings.ReplaceAll(alarmTime, ":", "")
	log.Printf("launch job[%s], alarmTime: %s, codes: %s\n", jobId, alarmTime, codes)

	t := getNextTick(alarmTime, jobId)
	for true {
		select {
		case <-t:
			log.Printf("job[%s] tick invoked\n", jobId)
			doProcess(codes, jobId)
			t = getNextTick(alarmTime, jobId)
		}
	}
}

func doProcess(codes []string, jobId string) {
	url := queryUrlPrefix + strings.Join(codes, ",")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("job[%s] get stock data error: %s", jobId, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := ConvertByte2String(body, GB18030)
	log.Printf("job[%s] get stock row data:\n%s", jobId, bodyStr)
	//parse body
	stocks := ParseStocks(bodyStr, codes)
	log.Println("get stocks: ", stocks)
	//push
	NewMessage(stocks).Push()
}

func getNextTick(alarmTime, jobId string) <-chan time.Time {
	now := time.Now()
	ts := strings.Split(alarmTime, ":")
	if len(ts) != 2 {
		log.Fatalf("alarmTime %s is invalid", alarmTime)
	}
	th, tm := ParseInt(ts[0]), ParseInt(ts[1])
	tt := time.Date(now.Year(), now.Month(), now.Day(), th, tm, 0, 0, ctsZone)
	log.Printf("job[%s] next tick time: %s", jobId, tt)
	d := tt.Sub(now)
	if d < 0 {
		d = tt.Add(time.Hour * 24).Sub(now)
	}
	log.Printf("job[%s] get next tick duration: %s", jobId, d)
	return time.Tick(d)
}
