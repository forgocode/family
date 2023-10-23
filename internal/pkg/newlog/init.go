package newlog

import (
	"go.uber.org/zap/zapcore"

	"github.com/forgocode/family/pkg/log"
)

var Logger *log.NewLogger

func InitLogger(fileName string, w zapcore.WriteSyncer) {
	opt := log.Options{
		FileName:   fileName,
		Level:      "info",
		ModuleName: "",
		W:          w,
	}
	Logger = log.New(opt)
}
