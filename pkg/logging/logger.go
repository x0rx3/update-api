package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logLevel string) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var lv zapcore.Level
	switch logLevel {
	case "debug":
		lv = zap.DebugLevel
	case "info":
		lv = zap.InfoLevel
	case "warn":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(lv),
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build())
}
