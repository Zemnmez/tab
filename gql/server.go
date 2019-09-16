package gql

import (
	"net/http"

	"github.com/zemnmez/tab/gql/resolver"
	"github.com/zemnmez/tab/storage"
	"github.com/zemnmez/tab/types"

	"github.com/99designs/gqlgen/handler"
)

type Server struct {
	// stupid bullshit goes here 😎
	Playground bool
	mux        *http.ServeMux
	resolver.Context
}

func (s Server) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	rq = s.Context.WithContext(rq)
	s.mux.ServeHTTP(rw, rq)
}

func NewServer(storage storage.IStorage) (s Server) {
	s.Context.Storage = storage.Storage{
		IStorage: storage,
		Tables: storage.NewTables(
			types.TableName{id: types.TABLE_DEFAULT},
			types.TableName{id: types.TABLE_LINKS},
		),
	}

	s.mux = http.NewServeMux()

	s.mux.Handle("/", handler.Playground("tab", "/query"))
	s.mux.Handle("/query", handler.GraphQL(Config{
		Resolvers: resolver.Resolver{},
	}))

	return
}
