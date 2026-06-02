package logx

import (
	"context"
	"log/slog"

	"github.com/samber/lo"
	"github.com/samber/oops"
)

type errorFormattingHandler struct {
	slog.Handler
}

func newErrorFormattingHandler(handler slog.Handler) slog.Handler {
	if handler == nil {
		return nil
	}
	return &errorFormattingHandler{Handler: handler}
}

func (h *errorFormattingHandler) Handle(ctx context.Context, record slog.Record) error {
	next := slog.NewRecord(record.Time, record.Level, record.Message, record.PC)
	record.Attrs(func(attr slog.Attr) bool {
		next.AddAttrs(normalizeErrorAttr(attr))
		return true
	})
	return h.Handler.Handle(ctx, next)
}

func (h *errorFormattingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := lo.Map(attrs, func(attr slog.Attr, _ int) slog.Attr {
		return normalizeErrorAttr(attr)
	})
	return &errorFormattingHandler{Handler: h.Handler.WithAttrs(next)}
}

func (h *errorFormattingHandler) WithGroup(name string) slog.Handler {
	return &errorFormattingHandler{Handler: h.Handler.WithGroup(name)}
}

func normalizeErrorAttr(attr slog.Attr) slog.Attr {
	if isErrorKey(attr.Key) {
		err, ok := errorFromValue(attr.Value)
		if ok {
			oopsErr, ok := oops.AsOops(err)
			if ok {
				attr.Value = oopsErr.LogValue()
				return attr
			}
		}
	}

	value := attr.Value.Resolve()
	if value.Kind() == slog.KindGroup {
		next := lo.Map(value.Group(), func(groupAttr slog.Attr, _ int) slog.Attr {
			return normalizeErrorAttr(groupAttr)
		})
		attr.Value = slog.GroupValue(next...)
		return attr
	}

	return attr
}

func isErrorKey(key string) bool {
	return key == "error" || key == "err"
}

func errorFromValue(value slog.Value) (error, bool) {
	if err, ok := value.Any().(error); ok {
		return err, true
	}

	resolved := value.Resolve()
	if err, ok := resolved.Any().(error); ok {
		return err, true
	}

	return nil, false
}
