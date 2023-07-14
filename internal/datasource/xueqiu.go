package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
)

const XUEQIU_API = ""

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) SupportBatchFetch() bool {
	return false
}

func (xq XueqiuDataSource) FetchData(code string) core.Stock {
	log.Sugar().Infof("FetchData")
	return core.Stock{}
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
