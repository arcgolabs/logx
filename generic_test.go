package logx_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	collectionmapping "github.com/arcgolabs/collectionx/mapping"
	"github.com/arcgolabs/logx"
)

func TestEnricher_GenericFieldMethods(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, nil))
	fields := collectionmapping.NewMapFrom(map[string]int{"attempt": 2})
	derived := logx.Enrich(logger).WithField("tenant", "acme")
	derived = logx.Enrich(derived).WithFields(fields)
	derived.Info("request")

	payload := map[string]any{}
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if got := payload["tenant"]; got != "acme" {
		t.Fatalf("tenant = %#v, want acme", got)
	}
	if got := payload["attempt"]; got != float64(2) {
		t.Fatalf("attempt = %#v, want 2", got)
	}
}

func TestEnricher_NilAndEmptyInputs(t *testing.T) {
	t.Parallel()

	if got := logx.Enrich(nil).WithField("retry", 3); got != nil {
		t.Fatalf("WithField() = %v, want nil", got)
	}

	logger := slog.New(slog.DiscardHandler)
	if got := logx.Enrich(logger).WithFields[int](nil); got != logger {
		t.Fatal("WithFields(nil) should return the original logger")
	}
}
