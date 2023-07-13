package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
	"sns/utils"
)

type StockDataFetcher interface {
	FetchData(code string) core.Stock
}

type StockBatchDataFetcher interface {
	FetchBatchData(codes []string) []core.Stock
}

type StockDataSource interface{}

var stockDataSources = make(map[string]StockDataSource)

func registerStockDataSource(name string, datasource StockDataSource) {
	if _, ok := datasource.(StockDataFetcher); !ok {
		log.Sugar().Errorf("try to add a non stockDataSource, name=%s, datasource=%s", name, datasource)
		return
	}
	if typeS := utils.TypeOf(datasource); typeS != "StockDataFetcher" && typeS != "StockBatchDataFetcher" {
		log.Sugar().Errorf("try to add a non stockDataSource, name=%s, datasource=%s", name, datasource)
		return
	}
	if stockDataSources[name] != nil {
		log.Sugar().Warnf("try to add stockDataSource with same name. name=%s, datasource=%s", name, datasource)
		return
	}
	stockDataSources[name] = datasource
	log.Sugar().Infof("add stockDataSource, name=%s", name)
}
