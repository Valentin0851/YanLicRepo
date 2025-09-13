package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	Key       = "logger"
	RequestId = "request_id"
)

var logger *zap.Logger

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, Key, &Logger{l: logger})

	return ctx, nil
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}
	l.l.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}
	l.l.Error(msg, fields...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}
	l.l.Debug(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}
	l.l.Warn(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}
	l.l.Fatal(msg, fields...)
}
