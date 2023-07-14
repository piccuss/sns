package datasource

import "sns/internal/pkg"

const XUEQIU_API = ""

type XueqiuDataSource struct {
}

func (xq XueqiuDataSource) supportBatchFetch() bool {
	return false
}

func (xq XueqiuDataSource) fetchData(code string) pkg.Stock {
	pkg.Sugar().Infof("FetchData")
	//TODO implement real method
	return pkg.Stock{}
}

func (xq XueqiuDataSource) fetchBatchData(codes []string) []pkg.Stock {
	pkg.Sugar().Infof("FetchBatchData")
	//TODO implement real method
	return []pkg.Stock{}
}

func init() {
	registerStockDataSource("xueqiu", XueqiuDataSource{})
}
