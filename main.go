package main

import (
	"go.uber.org/zap"

	"time"
)

func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./zap.log",
		"stderr",
		//"stdout",
	}
	return cfg.Build()
}
func main() {
	logger, _ := NewZapLogger()
	//zap.NewDevelopmentConfig()

	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "https://bilibili.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

}
