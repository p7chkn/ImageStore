package logger

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}
