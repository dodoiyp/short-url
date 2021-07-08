package global

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	now := time.Now()
	filename := fmt.Sprintf("%s/%04d-%02d-%02d.log", Conf.Logs.Path, now.Year(), now.Month(), now.Day())
	hook := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    Conf.Logs.MaxSize,
		MaxBackups: Conf.Logs.MaxBackups,
		MaxAge:     Conf.Logs.MaxAge,
		Compress:   Conf.Logs.Compress,
	}
	defer hook.Close()
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = ZapLogLocalTimeEncoder
	enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		Conf.Logs.Level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log = logger.Sugar()
	Log.Debug("init log done.")
}

func ZapLogLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(MsecLocalTimeFormat))
}
