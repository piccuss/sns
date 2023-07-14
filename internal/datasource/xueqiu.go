package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
)

const XUEQIU_API = ""

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) supportBatchFetch() bool {
	return false
}

func (xq XueqiuDataSource) fetchData(code string) core.Stock {
	log.Sugar().Infof("FetchData")
	//TODO implement real method
	return core.Stock{}
}

func (xq XueqiuDataSource) fetchBatchData(codes []string) []core.Stock {
	log.Sugar().Infof("FetchBatchData")
	//TODO implement real method
	return []core.Stock{}
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
