package gql

import (
	"net/http"
	"github.com/zemnmez/tab/gql/resolver"
	"github.com/zemnmez/tab/gql/storage"

	"github.com/99designs/gqlgen/handler"
)

type Server struct {
	// stupid bullshit goes here 😎
	Playground bool
	mux        *http.ServeMux
	resolver.ExecutionContext
}

func (s Server) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	rq = s.ExecutionContext.WithContext(rq)
	s.mux.ServeHTTP(rw, rq)
}

func NewServer(storage storage.IStorage) (s Server) {
	s.ExecutionContext.Storage = storage.Storage { storage }

	s.mux = http.NewServeMux()

	s.mux.Handle("/", handler.Playground("tab", "/query"))
	s.mux.Handle("/query", handler.GraphQL(Config{
		Resolvers: resolver.Resolver{},
	}))

	return
}
