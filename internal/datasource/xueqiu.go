package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
)

const XUEQIU_API = ""

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) StockDataFetcher() core.Stock {
	log.Sugar().Infof("StockDataFetcher")
	return core.Stock{}
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
