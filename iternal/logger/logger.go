package logger

import (
	"log"

	"go.uber.org/zap"
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
