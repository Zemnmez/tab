package tab // import "github.com/zemnmez/tab"

//go:generate go run github.com/99designs/gqlgen

import (
	uuid "github.com/satori/go.uuid"
)

type ID uuid.UUID
