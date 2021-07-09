package logger

import (
	"encoding/json"
	"fmt"
	"short-url/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var Log *zap.SugaredLogger

func InitLog(logConfig *config.LogsConfiguration) error {
	var err error = nil
	Logger, Log, err = InitZapLog(logConfig.Path, logConfig.Level)
	if err != nil {
		fmt.Println("Init lottery draw zap log is err: ", err)
		return fmt.Errorf("Init lottery draw zap log is err: %v", err)
	}

	return nil
}

/************************************
@ path: the path of log file
@ level: log level(DEBUG, INFO, ERROR)
@ return: *zap.Logger:  critical log
@ *zap.SugaredLogger: not critical log
@***********************************/
func InitZapLog(path string, level string) (*zap.Logger, *zap.SugaredLogger, error) {
	if path == "" || level == "" {
		fmt.Println("log path is nil or log level is nil")
		return nil, nil, fmt.Errorf("log path is nil or log level is nil")
	}

	jasonString := fmt.Sprintf(`{
             "level": "%s",
             "encoding": "json",
             "outputPaths": ["%s"],
             "errorOutputPaths": ["%s"]
             }`, level, path, path)

	var cfg zap.Config
	if err := json.Unmarshal([]byte(jasonString), &cfg); err != nil {
		return nil, nil, fmt.Errorf("init log json is err: %v", err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	Logger, err := cfg.Build()
	if err != nil {
		return nil, nil, fmt.Errorf("logger build error: %v", err)
	}

	Log := Logger.Sugar()

	return Logger, Log, nil
}
