package datasource

import "sns/internal/pkg"

const XUEQIU_API = ""

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) supportBatchFetch() bool {
	return false
}

func (xq XueqiuDataSource) fetchData(code string) (pkg.Stock, error) {
	pkg.Sugar().Infof("FetchData")
	//TODO implement real method
	return pkg.Stock{}, nil
}

func (xq XueqiuDataSource) fetchBatchData(codes []string) ([]pkg.Stock, error) {
	pkg.Sugar().Infof("FetchBatchData")
	//TODO implement real method
	return []pkg.Stock{}, nil
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
