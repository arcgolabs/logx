package logx

import (
	"bytes"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

const (
	consoleStacktraceKey = "_logx_stacktrace"
	consoleSourcesKey    = "_logx_sources"
)

type consoleMultilineField struct {
	key   string
	label string
}

var (
	consoleErrorKeys       = []string{"error", "err"}
	consoleMultilineFields = []consoleMultilineField{
		{key: consoleStacktraceKey, label: "stacktrace"},
		{key: consoleSourcesKey, label: "sources"},
	}
)

func newConsoleWriter(cfg config) zerolog.ConsoleWriter {
	writer := zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    cfg.timeFormat,
		NoColor:       cfg.noColor,
		FieldsExclude: []string{consoleStacktraceKey, consoleSourcesKey},

		FormatPrepare: prepareConsoleOopsFields,
		FormatExtra:   formatConsoleOopsFields,
	}
	return writer
}

func prepareConsoleOopsFields(evt map[string]any) error {
	lo.ForEach(consoleErrorKeys, func(key string, _ int) {
		extractConsoleOopsField(evt, key)
	})
	return nil
}

func extractConsoleOopsField(evt map[string]any, key string) {
	raw, ok := evt[key]
	if !ok {
		return
	}

	errorPayload, ok := raw.(map[string]any)
	if !ok {
		return
	}

	if stacktrace, ok := errorPayload["stacktrace"].(string); ok && stacktrace != "" {
		evt[consoleStacktraceKey] = stacktrace
		delete(errorPayload, "stacktrace")
	}

	if sources, ok := errorPayload["sources"].(string); ok && sources != "" {
		evt[consoleSourcesKey] = sources
		delete(errorPayload, "sources")
	}
}

func formatConsoleOopsFields(evt map[string]any, buf *bytes.Buffer) error {
	lo.ForEach(consoleMultilineFields, func(field consoleMultilineField, _ int) {
		writeConsoleMultilineField(evt, buf, field.key, field.label)
	})
	return nil
}

func writeConsoleMultilineField(evt map[string]any, buf *bytes.Buffer, key, label string) {
	value, ok := evt[key].(string)
	if !ok || value == "" {
		return
	}

	buf.WriteByte('\n')
	buf.WriteString(label)
	buf.WriteString(":\n")
	buf.WriteString(value)
}
