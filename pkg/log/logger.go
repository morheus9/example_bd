package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.Logger
	atomicLevel   zap.AtomicLevel
)

func Logger() *zap.Logger {
	return defaultLogger
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return defaultLogger.Check(lvl, msg)
}

func SetLevel(lvl zapcore.Level) {
	atomicLevel.SetLevel(lvl)
}

func Level() zapcore.Level {
	return defaultLogger.Level()
}

func WithLazy(fields ...zap.Field) *zap.Logger {
	return defaultLogger.WithLazy(fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return defaultLogger.With(fields...)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return defaultLogger.WithOptions()
}

func Named(s string) *zap.Logger {
	return defaultLogger.Named(s)
}

func Sugar() *zap.SugaredLogger {
	return defaultLogger.Sugar()
}

func Sync() error {
	return defaultLogger.Sync()
}

func Core() zapcore.Core {
	return defaultLogger.Core()
}

func Name() string {
	return defaultLogger.Name()
}

func Fatal(msg string, fields ...zap.Field) {
	defaultLogger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...any) {
	defaultLogger.Sugar().Fatalf(template, args...)
}

func Panic(msg string, fields ...zap.Field) {
	defaultLogger.Panic(msg, fields...)
}

func Panicf(template string, args ...any) {
	defaultLogger.Sugar().Panicf(template, args...)
}

func DPanic(msg string, fields ...zap.Field) {
	defaultLogger.DPanic(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	defaultLogger.Debug(msg, fields...)
}

func Debugf(template string, args ...any) {
	defaultLogger.Sugar().Debugf(template, args...)
}

func Info(msg string, fields ...zap.Field) {
	defaultLogger.Info(msg, fields...)
}

func Infof(template string, args ...any) {
	defaultLogger.Sugar().Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	defaultLogger.Warn(msg, fields...)
}

func Warnf(template string, args ...any) {
	defaultLogger.Sugar().Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	defaultLogger.Error(msg, fields...)
}

func Errorf(template string, args ...any) {
	defaultLogger.Sugar().Errorf(template, args...)
}

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	defaultLogger.Log(lvl, msg, fields...)
}
