package gql

import (
	"context"
	"net/http"
)

type authorizationHeaderKeyT struct{}

var authorizationHeaderKey = &authorizationHeaderKeyT{}

func (Server) authzContext(rq *http.Request) (withContext *http.Request) {
	return rq.WithContext(
		context.WithValue(
			rq.Context(),
			authorizationHeaderKey,
			rq.Header.Get("Authorization"),
		),
	)
}
