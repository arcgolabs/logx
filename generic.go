package logx

import (
	"log/slog"
)

// FieldSet is the minimal map-like shape accepted by Enricher.WithFields.
type FieldSet[V any] interface {
	Len() int
	Range(func(string, V) bool)
}

// Enricher derives slog loggers with strongly typed fields.
type Enricher struct {
	logger *slog.Logger
}

// Enrich creates a typed field enricher for logger.
func Enrich(logger *slog.Logger) Enricher {
	return Enricher{logger: logger}
}

// WithField adds one typed field and returns a derived logger.
func (e Enricher) WithField[V any](key string, value V) *slog.Logger {
	if e.logger == nil {
		return nil
	}
	return e.logger.With(key, value)
}

// WithFields adds typed fields and returns a derived logger.
func (e Enricher) WithFields[V any](fields FieldSet[V]) *slog.Logger {
	if e.logger == nil {
		return nil
	}
	if fields == nil || fields.Len() == 0 {
		return e.logger
	}

	args := make([]any, 0, fields.Len()*2)
	fields.Range(func(key string, value V) bool {
		args = append(args, key, value)
		return true
	})
	return e.logger.With(args...)
}
