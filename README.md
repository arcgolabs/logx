## Overview

`logx` is an opinionated logger builder that returns a standard `*slog.Logger`, backed by `zerolog`:

- You configure output (console / file + rotation), level, caller, and global logger options via `logx.With...`.
- You keep using the standard `slog` API (`Info`, `Error`, `With`, `WithGroup`...) in your application code.

## Install

```bash
go get github.com/arcgolabs/logx@latest
```

## Current capabilities

- `*slog.Logger` output backed by `zerolog`
- Console output and file output (+ rotation via `lumberjack`)
- Optional caller (`WithCaller(true)`) and optional global `zerolog` logger (`WithGlobalLogger()`)
- Trace/span fields from OpenTelemetry context (`WithTraceContext`)
- Structured `oops` error output through regular `slog` error fields

## Documentation map

- Minimal usage: [Getting Started](./docs/getting-started.md)
- Output / rotation / defaults: [Configuration](./docs/configuration.md)
- Trace context + oops: [Trace and oops](./docs/trace-and-oops.md)

## Runnable examples

- Trace context: [examples/logx/trace_context](./examples/trace_context/main.go)
- oops integration: [examples/oops_integration](./examples/oops_integration/main.go)
