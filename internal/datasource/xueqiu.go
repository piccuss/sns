package datasource

import (
	"errors"
	"sns/internal/pkg"
)

const XUEQIU_API = ""

const XUEQIU_CURL_DEMO = `curl 'https://stock.xueqiu.com/v5/stock/realtime/quotec.json?symbol=SH601231,SZ002299' \
-H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82' \
--compressed`

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) supportBatchFetch() bool {
	return true
}

func (xq XueqiuDataSource) fetchData(code string) (*pkg.Stock, error) {
	pkg.Sugar().Infof("xuqiu don't support FetchBatchData")
	return nil, errors.New("invoke unsupport method")
}

func (xq XueqiuDataSource) fetchBatchData(codes []string) ([]*pkg.Stock, error) {
	pkg.Sugar().Infof("xuqiu don't support FetchBatchData")
	//TODO implement BatchFetch
	return nil, errors.New("invoke unsupport method")
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
