package logs

import (
	"chatgpt-web-new-go/common/env"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var (
	LoggerZap *zap.Logger        // log
	Logger    *zap.SugaredLogger // log

	Debug func(template string, args ...interface{})
	Info  func(template string, args ...interface{})
	Warn  func(template string, args ...interface{})
	Error func(template string, args ...interface{})
)

func LogSyncLast() {
	func() {
		_ = Logger.Sync()
		_ = LoggerZap.Sync()
	}()
}

func LogInit() {
	hook := getWriter("./log/gpt.log")
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(hook)}

	// 如果是开发环境，同时在控制台上也输出
	if !env.IsProduction() {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	//field := zap.Fields(zap.String("env", IsProduction()))

	// 构造日志
	logger := zap.New(core, caller, development)
	LoggerZap = logger
	Logger = logger.Sugar()

	{
		Debug = Logger.Debugf
		Info = Logger.Infof
		Warn = Logger.Warnf
		Error = Logger.Errorf
	}

	Logger.Info("log initializer success")
}

func getWriter(filename string) io.Writer {
	// 生成rotateLogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// 保存90天内的日志，每1天(整点)分割一次日志
	hook, err := rotateLogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotateLogs.WithLinkName(filename),
		rotateLogs.WithMaxAge(time.Hour*24*365),
		rotateLogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
