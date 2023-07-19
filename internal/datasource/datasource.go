package datasource

import (
	"errors"
	"fmt"
	"sns/internal/pkg"
)

type StockDataSource interface {
	supportBatchFetch() bool
	fetchData(code string) (*pkg.Stock, error)
	fetchBatchData(codes []string) ([]*pkg.Stock, error)
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

func FetchStockData(name string, codes []string) ([]*pkg.Stock, error) {
	stockDataSource, ok := stockDataSources[name]

	if !ok {
		return nil, errors.New(fmt.Sprintf("cannot find dataSource=%s, please check config.", name))
	}

	if stockDataSource.supportBatchFetch() {
		return stockDataSource.fetchBatchData(codes)
	}
	//TODO implement aysnc get
	return []*pkg.Stock{}, nil
}
