---
title: 'logx Trace and oops'
linkTitle: 'trace-and-oops'
description: 'Attach trace/span IDs and work with oops errors'
weight: 4
---

## Trace context

If you use OpenTelemetry, you can attach trace/span IDs from a `context.Context`:

```go
package main

import (
	"context"

	"github.com/arcgolabs/logx"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	logger, err := logx.New(logx.WithConsole(true), logx.WithDebugLevel())
	if err != nil {
		panic(err)
	}
	defer func() { _ = logx.Close(logger) }()

	traceID, err := trace.TraceIDFromHex("4bf92f3577b34da6a3ce929d0e0e4736")
	if err != nil {
		panic(err)
	}
	spanID, err := trace.SpanIDFromHex("00f067aa0ba902b7")
	if err != nil {
		panic(err)
	}

	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceID,
		SpanID:     spanID,
		TraceFlags: trace.FlagsSampled,
	})
	ctx := trace.ContextWithSpanContext(context.Background(), spanContext)

	logx.WithTraceContext(ctx, logger).Info("request accepted", "endpoint", "/api/orders")
}
```

Runnable example:

- [examples/trace_context](../examples/trace_context/main.go)

## oops errors

`logx` formats `oops`-compatible errors through the standard `slog` API. Use
regular `error` or `err` fields; `logx` will keep the error payload structured.
JSON output keeps stacktrace line breaks as standard `\n` escapes, so log
processors can restore them as real newlines. Console output extracts the
stacktrace into a readable multi-line block.

```go
package main

import (
	"github.com/arcgolabs/logx"
	"github.com/samber/oops"
)

func main() {
	logger := logx.MustNew(logx.WithConsole(true))
	defer func() { _ = logx.Close(logger) }()

	err := oops.In("payment").Errorf("upstream failed")
	logger.Error("request failed", "error", err)
}
```

Runnable example:

- [examples/oops_integration](../examples/oops_integration/main.go)
