package gql

import (
	"net/http"
	"github.com/zemnmez/tab/gql/resolver"

	"github.com/99designs/gqlgen/handler"
)

type Server struct {
	// stupid bullshit goes here 😎
	Playground bool
	mux        *http.ServeMux
}

func (s Server) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	s.mux.ServeHTTP(rw, s.authzContext(rq))
}

func NewServer() (s Server) {
	s.mux = http.NewServeMux()

	s.mux.Handle("/", handler.Playground("tab", "/query"))
	s.mux.Handle("/query", handler.GraphQL(Config{
		Resolvers: resolver.Resolver{},
	}))

	return
}
