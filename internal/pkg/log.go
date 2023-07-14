package pkg

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func init() {
	zlog, _ := zap.NewProduction()
	sugar = zlog.Sugar()
}

func Sugar() *zap.SugaredLogger {
	return sugar
}
