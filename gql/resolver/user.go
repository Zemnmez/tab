package resolver

import (
	"context"
	"io"

	"github.com/zemnmez/tab/types"
)

type SelfMutation struct{}

type AnonymousUser struct{}

type Self struct {
	AuthenticationTokenInfo
}

type UserQuery struct{}
type UserMutation struct{}

func (UserQuery) Self(ctx context.Context) Self {
	auth := Context.Get(nil, ctx).Authentication
	return Self { AuthenticationQuery{}.Token(Auth) }
}

func (UserQuery) WhoCan(do []types.Authorization) (users []User, err error) { panic("todo") }

func (UserMutation) Self() UserMutator    { panic("todo") }
func (UserMutation) Special() UserMutator { panic("todo") }
func (UserMutation) Regular() UserMutator { panic("todo") }

func (Mutation) User() UserMutation { return UserMutation{} }

type UserMutator interface {
	Modify(with UserInput) User
}

type UserID interface {
	io.WriterTo
	io.ReaderFrom
}

type User interface {
	Name() string
	Authentication() *UserAuthentication
	Grants() []AuthorizationGrant
	Grant(u User)
	History() []HistoryItem
}

// Self can be any User.
type Self struct {
	User
}

type RegularUser struct{ types.RegularUser }

func (r RegularUser) Authentication() UserAuthentication { panic("todo") }
func (s RegularUser) History() []HistoryItem             { panic("todo") }
func (r RegularUser) Grants() []AuthorizationGrant       { panic("todo") }

type RootUser struct{ types.User }
