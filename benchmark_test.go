package logx_test

import (
	"log/slog"
	"testing"

	collectionmapping "github.com/arcgolabs/collectionx/mapping"
	"github.com/arcgolabs/logx"
)

func benchmarkLogger(b *testing.B) *slog.Logger {
	b.Helper()

	return newTestLogger(b, logx.WithConsole(false), logx.WithDebugLevel())
}

func BenchmarkLoggerInfo(b *testing.B) {
	logger := benchmarkLogger(b)

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		logger.Info("benchmark message", "key", "value", "count", 1)
	}
}

func BenchmarkLoggerWithFieldsInfo(b *testing.B) {
	logger := benchmarkLogger(b)
	fields := collectionmapping.NewMapFrom(map[string]any{
		"service": "arcgo",
		"env":     "bench",
	})

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		logx.Enrich(logger).WithFields(fields).Info("with-fields")
	}
}

func BenchmarkSlogInfo(b *testing.B) {
	slogLogger := benchmarkLogger(b)

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		slogLogger.Info("slog benchmark", "key", "value", "count", 1)
	}
}
