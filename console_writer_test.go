package logx

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestConsoleWriterFormatsOopsStacktraceAsMultiline(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer(nil)
	writer := newConsoleWriter(config{
		noColor:    true,
		timeFormat: "2006-01-02 15:04:05",
	})
	writer.Out = buffer

	logger := zerolog.New(writer)
	logger.Error().
		Interface("error", map[string]any{
			"err":        "payment provider failed",
			"stacktrace": "Oops: payment provider failed\n  --- at payment.go:42",
		}).
		Msg("checkout failed")

	output := buffer.String()
	if strings.Contains(output, `stacktrace":"Oops: payment provider failed\n`) {
		t.Fatalf("expected stacktrace to be extracted from inline error JSON, got %q", output)
	}
	if !strings.Contains(output, "stacktrace:\nOops: payment provider failed\n  --- at payment.go:42") {
		t.Fatalf("expected multiline stacktrace block, got %q", output)
	}
	if strings.Count(output, "Oops: payment provider failed") != 1 {
		t.Fatalf("expected stacktrace to be printed once, got %q", output)
	}
}
