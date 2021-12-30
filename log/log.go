package logger

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go-kratos/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger(c *conf.Config) *Logger {
	return &Logger{log: newZapLogger(c)}
}

func newZapLogger(c *conf.Config) *zap.Logger {

	infoWriteSyncer := initInfoLogWriter(c)
	errWriteSyncer := initErrLogWriter(c)
	encoder := getEncoder()

	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.InfoLevel
	})

	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.InfoLevel && lev >= zap.DebugLevel
	})

	errPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoWriteSyncer, infoPriority),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdin), debugPriority),
		zapcore.NewCore(encoder, errWriteSyncer, errPriority),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 根据上面的配置创建logger
	return logger
	// zap.ReplaceGlobals(logger)               // 替换zap库里全局的logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig) // json格式日志
}

func initInfoLogWriter(c *conf.Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.Log.FileName,
		MaxSize:    c.Log.MaxSize,    // 日志文件大小 单位：MB
		MaxBackups: c.Log.MaxBackups, // 备份数量
		MaxAge:     c.Log.MaxAge,     // 备份时间 单位：天
		Compress:   true,             // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func initErrLogWriter(c *conf.Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.LogErr.FileName,
		MaxSize:    c.LogErr.MaxSize,    // 日志文件大小 单位：MB
		MaxBackups: c.LogErr.MaxBackups, // 备份数量
		MaxAge:     c.LogErr.MaxAge,     // 备份时间 单位：天
		Compress:   true,                // 是否压缩
	}

	return zapcore.AddSync(lumberJackLogger)
}

func (l *Logger) Log(level log.Level, keyVal ...interface{}) error {
	args := l.argsToFile(keyVal)

	switch level {
	case log.LevelDebug:
		l.log.Debug("", args...)
	case log.LevelInfo:
		l.log.Info("", args...)
	case log.LevelWarn:
		l.log.Warn("", args...)
	case log.LevelError:
		l.log.Error("", args...)
	}
	return nil
}

func (l *Logger) argsToFile(args ...interface{}) []zap.Field {
	res := args[0].([]interface{})
	if len(res) <= 0 {
		return nil
	}

	if len(res)%2 != 0 {
		l.log.Error(fmt.Sprint("args must appear in pairs: ", args))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(res); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(res[i]), fmt.Sprint(res[i+1])))
	}
	return data
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log.Debug(msg, l.argsToFile(args)...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log.Info(msg, l.argsToFile(args)...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log.Error(msg, l.argsToFile(args)...)
}
