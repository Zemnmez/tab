package resolver

type resolverContextKey struct{}
var (
	contextKey = &contextKey{}
)

type Context struct {
	storage.Storage
	Authentication string
}

func (e *ExecutionContext) Get(ctx context.Context) ExecutionContext { return ctx.Value(contextKey).(ExecutionContext) }

func (e ExecutionContext) WithContext(rq *http.Request) *http.Request {
	e.Authentication = rq.Header.Get("Authentication")

	ctx := rq.Context()

	ctx = context.WithValue(
		ctx,
		contextKey,
		e,
	)

	return rq.WithContext(ctx)
}