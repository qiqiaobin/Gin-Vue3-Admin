package logger

import (
	"log"
	"os"
	"strings"
	"tadmin/conf"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 初始化日志
/*初始化日志
 * filename 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 * 由于zap不具备日志切割功能, 这里使用lumberjack配合
 */
func InitLogger() {
	var logCore zapcore.Core

	logLevel, ok := levelStrToZapLevel[strings.ToLower(conf.Config.Logger.Level)]
	if !ok {
		log.Fatalln("日志级别设置错误")
	}

	if conf.Config.Logger.Filename == "" {
		conf.Config.Logger.Filename = "./logs/out.log" // 默认日志文件
	}

	// writers
	writersSyncers := GetLoggerWriter()

	// zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	// 但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	enConfig := zap.NewProductionEncoderConfig() // 生成配置

	// 时间格式
	enConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		//encoder.AppendString(t.Format("2006-01-02 15:04:05.999"))
		encoder.AppendString(t.Format(time.DateTime))
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 判断程序当前所处的模式，开发模式所有的日志打印到控制台即可
	if conf.Config.System.Env == "dev" {
		// 将多个 zapcore.Core 对象合并成一个
		logCore = zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		)
	} else {
		logCore = zapcore.NewCore(
			zapcore.NewConsoleEncoder(enConfig), // 编码器配置
			writersSyncers,                      // 打印文件
			logLevel,                            // 日志等级
		)
	}

	// 开启开发模式，堆栈跟踪
	if conf.Config.Logger.ShowLine {
		caller := zap.AddCaller()
		// 构造日志
		logger := zap.New(logCore, caller, zap.AddCallerSkip(1))
		sugar = logger.Sugar()
	} else {
		// 构造日志
		logger := zap.New(logCore)
		sugar = logger.Sugar()
	}

	sugar.Info("初始化日志完成")
}

// GetLoggerWriter return writerSyncer
// @param *settings.Log config 日志配置信息
// @return zapcore.WriteSyncer 返回一个日志记录器
func GetLoggerWriter() zapcore.WriteSyncer {
	lumberLoggers := &lumberjack.Logger{
		Filename:   conf.Config.Logger.Filename,
		MaxSize:    conf.Config.Logger.MaxSize,
		MaxBackups: conf.Config.Logger.MaxBackups,
		MaxAge:     conf.Config.Logger.MaxAge,
		LocalTime:  true, //backup的日志是否使用本地时间戳，默认使用UTC时间
		Compress:   conf.Config.Logger.Compress,
	}
	defer lumberLoggers.Close()
	// zapcore.Lock() 确保多个 goroutine 同时写入日志时不会出现并发问题
	return zapcore.Lock(zapcore.AddSync(lumberLoggers))
}

// GetEncoder 将日志消息编码为指定的格式
// @return zapcore.Encoder 返回Encoder
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.FunctionKey = "function"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
