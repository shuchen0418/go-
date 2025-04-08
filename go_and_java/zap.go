package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
)

func main() {
	InitLogger()
	defer logger.Sync()
	logger.Info("hello world",
		zap.String("sb", "hz"),
		zap.String("sss", "sss"))
}

func InitLogger() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("D:/Desktop/go 八股文/test.log")
	return zapcore.AddSync(file)
}
