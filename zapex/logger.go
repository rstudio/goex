package zapex

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const (
	loggerCtxKey = contextKey("zapex.logger")
)

type contextKey string

// Creates a zap logger or dies trying. (No, seriously, it will log.Fatal() if it fails).
// The caller is responsible for calling `defer logger.Sync()` after initializing.
func NewLogger(envName string, debug bool) *zap.Logger {
	if envName == "dev" || envName == "test" {
		zCfg := zap.NewDevelopmentConfig()
		zCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zl, err := zCfg.Build()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}
		return zl
	}

	// Customize the encoder for datadog-friendliness
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder

	logLevel := zap.InfoLevel
	if debug {
		logLevel = zap.DebugLevel
	}

	// Modified from zap.NewProduction
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	zl, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return zl
}

func ContextWithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, logger)
}

func LoggerFromContext(ctx context.Context) (*zap.SugaredLogger, bool) {
	logger, ok := ctx.Value(loggerCtxKey).(*zap.SugaredLogger)
	return logger, ok
}

func LoggerWithTraceFields(ctx context.Context, logger *zap.SugaredLogger) *zap.SugaredLogger {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		logger.Debugw("no tracing span found in context", "ctx", ctx)
		return logger
	}

	spanCtx := span.Context()

	traceID := fmt.Sprintf("%v", spanCtx.TraceID())
	spanID := fmt.Sprintf("%v", spanCtx.SpanID())

	logger.Debugw(
		"returning logger with trace fields",
		"dd.trace_id", traceID,
		"dd.span_id", spanID,
	)

	return logger.With(
		"dd.trace_id", traceID,
		"dd.span_id", spanID,
	)
}

func NewStdlibWrappedLogger(ctx context.Context, logger *zap.SugaredLogger) *log.Logger {
	pr, pw := io.Pipe()
	scanner := bufio.NewScanner(pr)
	logger = logger.With("source", "stdlib")

	go func() {
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
				logger.Info(scanner.Text())
			}
		}
	}()

	return log.New(pw, "", 0)
}
