---
title: 'logx Getting Started'
linkTitle: 'getting-started'
description: 'Create a slog logger, log with fields, and close resources safely'
weight: 2
---

## Getting started

`logx.New(...)` returns a standard `*slog.Logger`. The module requires Go 1.27 or newer. If you enable file output, make sure to close resources via `logx.Close(logger)` (it is safe to call even when only console output is used).

## Minimal example

```go
package main

import (
	"log/slog"

	"github.com/arcgolabs/logx"
)

func main() {
	logger, err := logx.New(
		logx.WithConsole(true),
		logx.WithLevel(slog.LevelInfo),
		logx.WithCaller(true),
	)
	if err != nil {
		panic(err)
	}
	defer func() { _ = logx.Close(logger) }()

	logger.Info("service started", "service", "user-api")

	reqLogger := logx.Enrich(logger).WithField("request_id", "req_123")
	reqLogger.Info("request accepted", "path", "/api/health")
}
```

`Enricher.WithField` and `Enricher.WithFields` are Go 1.27 generic methods. `WithFields` accepts map-like collections whose values have any concrete type, including `collectionx/mapping.Map`.

## Next

- Output / rotation / defaults: [Configuration](./configuration.md)
- Trace/span context and oops: [Trace and oops](./trace-and-oops.md)
