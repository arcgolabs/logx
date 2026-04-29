package logx

import (
	"log/slog"

	collectionmapping "github.com/arcgolabs/collectionx/mapping"
)

// WithFieldT adds one typed field to logger and returns derived logger.
func WithFieldT[T any](logger *slog.Logger, key string, value T) *slog.Logger {
	if logger == nil {
		return nil
	}
	return WithField(logger, key, value)
}

// WithFieldsT adds typed fields to logger and returns derived logger.
func WithFieldsT[T any](logger *slog.Logger, fields Fields[T]) *slog.Logger {
	if logger == nil {
		return nil
	}
	converted := collectionmapping.NewMapWithCapacity[string, any](fields.Len())
	fields.Range(func(key string, value T) bool {
		converted.Set(key, value)
		return true
	})
	return WithFields(logger, converted)
}
