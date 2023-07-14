package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
)

type StockDataSource interface {
	supportBatchFetch() bool
	fetchData(code string) core.Stock
	fetchBatchData(codes []string) []core.Stock
}

var stockDataSources = make(map[string]StockDataSource)

func registerStockDataSource(name string, datasource StockDataSource) {
	if stockDataSources[name] != nil {
		log.Sugar().Warnf("try to add stockDataSource with same name. name=%s, datasource=%s", name, datasource)
		return
	}
	stockDataSources[name] = datasource
	log.Sugar().Infof("add stockDataSource, name=%s", name)
}

func GetStockDataSource(name string) StockDataSource {
	return stockDataSources[name]
}

func FetchStockData(codes []string, stockDataSource StockDataSource) []core.Stock {
	if stockDataSource.supportBatchFetch() {
		return stockDataSource.fetchBatchData(codes)
	}
	//TODO implement aysnc get
	return []core.Stock{}
}
