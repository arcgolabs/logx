// Package main demonstrates oops error logging through standard slog fields.
package main

import (
	"errors"
	"log/slog"
	"time"

	"github.com/arcgolabs/logx"
	"github.com/samber/oops"
)

func main() {
	logger := logx.MustNew(
		logx.WithConsole(true),
		logx.WithDebugLevel(),
	)
	defer func() {
		if err := logx.Close(logger); err != nil {
			panic(err)
		}
	}()

	logPaymentFailure(logger)
	logWrappedFailure(logger)
	logDerivedLoggerFailure(logger)
}

func logPaymentFailure(logger *slog.Logger) {
	err := oops.
		In("payment").
		Code("payment.provider_unavailable").
		Tags("checkout", "retryable").
		With("provider", "stripe").
		With("attempt", 2).
		Since(time.Now().Add(-150 * time.Millisecond)).
		Errorf("payment provider failed")

	logger.Error(
		"checkout failed",
		"order_id", "ord_1001",
		"error", err,
	)
}

func logWrappedFailure(logger *slog.Logger) {
	cause := errors.New("connection refused")
	err := oops.
		In("inventory").
		With("sku", "sku_42").
		Wrapf(cause, "reserve inventory")

	logger.Error(
		"inventory reservation failed",
		"err", err,
	)
}

func logDerivedLoggerFailure(logger *slog.Logger) {
	err := oops.
		User("user_123", "email", "buyer@example.com").
		Tenant("tenant_a").
		Hint("retry with a new idempotency key").
		New("idempotency key conflict")

	logger.
		With("request_id", "req_abc").
		Error("request rejected", "error", err)
}
