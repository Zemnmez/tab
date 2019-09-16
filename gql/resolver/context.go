package resolver

import (
	"context"
	"net/http"

	"github.com/zemnmez/tab/storage"
)

type resolverContextKey struct{}

var (
	contextKey = &resolverContextKey{}
)

type Context struct {
	storage.Storage
	Authorization AuthorizationToken
}

func (e *Context) Get(ctx context.Context) Context {
	return ctx.Value(contextKey).(Context)
}

func (e Context) WithContext(rq *http.Request) *http.Request {
	e.Authorization = AuthorizationToken(rq.Header.Get("Authorization"))

	ctx := rq.Context()

	ctx = context.WithValue(
		ctx,
		contextKey,
		e,
	)

	return rq.WithContext(ctx)
}
