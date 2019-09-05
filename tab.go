package tab // import "github.com/zemnmez/tab"

//go:generate go run github.com/99designs/gqlgen
//go:generate go run github.com/zemnmez/tab/proto $GOPATH proto/tab.proto

import (
	uuid "github.com/satori/go.uuid"
)

type ID uuid.UUID

type ID interface {
	ID() string
}
