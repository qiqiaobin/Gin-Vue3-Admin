package logger

import (
	"runtime/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	DPanicLevel = "dpanic"
	PanicLevel  = "panic"
	FatalLevel  = "fatal"
)

var levelStrToZapLevel = map[string]zapcore.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}

// zap日志
var sugar *zap.SugaredLogger

// Debug logger.Debug("test")
// Debug 格式化输出1个 debug 信息
func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

// Debugf logger.Debugf("test:%s", err)
// Debugf 格式化输出1个 debug 信息
func Debugf(format string, args ...interface{}) {
	sugar.Debugf(format, args...)
}

// Debugw logger.Debugw("test", "field1", "value1", "field2", "value2")
func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

// Info 格式化输出1个 info 信息
func Info(args ...interface{}) {
	sugar.Info(args...)
}

// Infof 格式化输出1个 info 信息
func Infof(format string, args ...interface{}) {
	sugar.Infof(format, args...)
}

// Infow 格式化输出1个 info 信息
func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

// Warn 格式化输出1个 Warn 信息
func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

// Warnf 格式化输出1个 Warn 信息
func Warnf(format string, args ...interface{}) {
	sugar.Warnf(format, args...)
}

// Warnw 格式化输出1个 Warn 信息
func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

// Error 输出一条 Error 信息
func Error(args ...interface{}) {
	sugar.Error(args...)
}

// Errorf 格式化输出1个 Error 信息
func Errorf(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}

// Errorw 格式化输出1个 Error 信息
func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

// Panic 输出一条 Panic 信息
func Panic(args ...interface{}) {
	sugar.Panic(args...)
}

// Panicf 格式化输出1个 panic 信息
func Panicf(format string, args ...interface{}) {
	sugar.Panicf(format, args...)
}

// Panicw 格式化输出1个 Error 信息
func Panicw(msg string, keysAndValues ...interface{}) {
	sugar.Panicw(msg, keysAndValues...)
}

// Fatal 输出一条 Fatal 信息
func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

// Fatalf 格式化输出1个 Fatal 信息
func Fatalf(format string, args ...interface{}) {
	debug.PrintStack()
	sugar.Fatalf(format, args...)
}

// Fatalw 格式化输出1个 Error 信息
func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}

// TraceError prints the stack and error
func TraceError(format string, args ...interface{}) {
	Error(string(debug.Stack()))
	Errorf(format, args...)
}
