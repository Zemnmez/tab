package user

import (
	"github.com/zemnmez/tab/generated"
)

// RegularID is the ID of a regular user
type RegularID struct{ uuid.Uuid }

// UserID implements users.ID
func (r RegularID) UserID() string { return r.String() }

type SpecialID = generated.SpecialUserID


type ID interface { UserID() string }

// A generic user
type User interface {
	Name() string
	Authentication() (UserAuthentication, error)
	Grants() ([]AuthorizationGrant, error)
	Authorizations() (allowedTo []Authorization, err error)
	History() ([]HistoryItem, error)
}

type Abstract struct {
	ID     UserID
	Grants []AuthorizationGrant
}

func (a Abstract) Authorizations() (allowedTo []Authorization, err error) { panic("unimplemented") }
func (a Abstract) History() (items []HistoryItem, err error)              { panic("unimplemented") }

type Special struct {
	ID SpecialID
	Abstract
}

func (s Special) Name() string { return s.ID }

type Regular struct {
	ID   RegularID
	Name string
	Abstract
}

func (r Regular) Name() string { return r.Name }
