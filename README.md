## Overview

`logx` is an opinionated logger builder that returns a standard `*slog.Logger`, backed by `zerolog`:

- You configure output (console / file + rotation), level, caller, and global logger options via `logx.With...`.
- You keep using the standard `slog` API (`Info`, `Error`, `With`, `WithGroup`...) in your application code.

## Install

```bash
go get github.com/arcgolabs/logx@latest
```

`logx` now requires Go 1.27 or newer.

## Current capabilities

- `*slog.Logger` output backed by `zerolog`
- Console output and file output (+ rotation via `lumberjack`)
- Optional caller (`WithCaller(true)`) and optional global `zerolog` logger (`WithGlobalLogger()`)
- Typed field enrichment through Go 1.27 generic methods (`logx.Enrich(logger).WithField/WithFields`)
- Trace/span fields from OpenTelemetry context (`WithTraceContext`)
- Structured `oops` error output through regular `slog` error fields

## Go 1.27 typed fields

```go
requestLogger := logx.Enrich(logger).WithField("request_id", "req_123")
requestLogger = logx.Enrich(requestLogger).WithFields(typedFields)
```

The package-level `WithField`, `WithFieldT`, `WithFields`, and `WithFieldsT` helpers were replaced by the generic `Enricher` methods above.

## Documentation map

- Minimal usage: [Getting Started](./docs/getting-started.md)
- Output / rotation / defaults: [Configuration](./docs/configuration.md)
- Trace context + oops: [Trace and oops](./docs/trace-and-oops.md)

## Runnable examples

- Trace context: [examples/logx/trace_context](./examples/trace_context/main.go)
- oops integration: [examples/oops_integration](./examples/oops_integration/main.go)
