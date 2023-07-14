package datasource

import "sns/internal/pkg"

type StockDataSource interface {
	supportBatchFetch() bool
	fetchData(code string) pkg.Stock
	fetchBatchData(codes []string) []pkg.Stock
}

var stockDataSources = make(map[string]StockDataSource)

func registerStockDataSource(name string, datasource StockDataSource) {
	if stockDataSources[name] != nil {
		pkg.Sugar().Warnf("try to add stockDataSource with same name. name=%s, datasource=%s", name, datasource)
		return
	}
	stockDataSources[name] = datasource
	pkg.Sugar().Infof("add stockDataSource, name=%s", name)
}

func GetStockDataSource(name string) StockDataSource {
	return stockDataSources[name]
}

func FetchStockData(codes []string, stockDataSource StockDataSource) []pkg.Stock {
	if stockDataSource.supportBatchFetch() {
		return stockDataSource.fetchBatchData(codes)
	}
	//TODO implement aysnc get
	return []pkg.Stock{}
}
