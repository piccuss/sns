package datasource

import (
	"sns/internal/core"
	"sns/internal/log"
)

type StockDataFetcher interface {
	FetchData(code string) core.Stock
}

type StockBatchDataFetcher interface {
	FetchBatchData(codes []string) []core.Stock
}

type StockDataSource interface {
	SupportBatchFetch() bool
}

var stockDataSources = make(map[string]StockDataSource)

func registerStockDataSource(name string, datasource StockDataSource) {
	if datasource.SupportBatchFetch() {
		if _, ok := datasource.(StockBatchDataFetcher); !ok {
			log.Sugar().Errorf("try to add a invalid stockDataSource, which supportBatchFetch but is not a StockBatchDataFetcher, name=%s, datasource=%s", name, datasource)
			return
		}
	} else {
		if _, ok := datasource.(StockDataFetcher); !ok {
			log.Sugar().Errorf("try to add a invalid stockDataSource, which not supportBatchFetch but is not a StockDataFetcher, name=%s, datasource=%s", name, datasource)
			return
		}
	}
	if stockDataSources[name] != nil {
		log.Sugar().Warnf("try to add stockDataSource with same name. name=%s, datasource=%s", name, datasource)
		return
	}
	stockDataSources[name] = datasource
	log.Sugar().Infof("add stockDataSource, name=%s", name)
}
