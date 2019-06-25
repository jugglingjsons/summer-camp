package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(SummerAppService) SummerAppService

type loggingMiddleware struct {
	logger log.Logger
	next   SummerAppService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SummerAppService) SummerAppService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) Create(ctx context.Context, payload string) (response string, err error) {
	defer func() {
		l.logger.Log("method", "Create", "payload", payload, "err", err)
	}()
	return l.next.Create(ctx, payload)
}

func (l loggingMiddleware) Health(ctx context.Context) (ok bool) {
	defer func() {
		l.logger.Log("method", "Health", "ok", ok)
	}()
	return l.next.Health(ctx)
}
