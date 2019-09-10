package resolver

import (
	"context"

	"github.com/zemnmez/tab/types"
)

type UserQuery struct{}
type UserMutation struct{}

func (UserQuery) Self(ctx context.Context) Self                             {
	auth := ExecutionContext.Get(nil, ctx).Authentication
	if auth == "" {

	}

}
func (UserQuery) Special(id types.SpecialUserID) SpecialUser          { panic("todo") }
func (UserQuery) Regular(id types.RegularUserID) RegularUser          { panic("todo") }
func (UserQuery) WhoCan(do []types.Authorization) (users []User, err error) { panic("todo") }

func (UserMutation) Self() UserMutator { panic("todo") }
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
	GrantSpecial(s SpecialUser)
	History() []HistoryItem
}

// Self can be any User.
type Self struct {
	User
}

type SpecialUser struct { types.SpecialUser }

var AnonymousUserDefault = SpecialUser {
	types.User {
		Id: types.ANONYMOUS,
		Authorizations: []Authorization { },
	}
}

func (s SpecialUser) Authentication() UserAuthentication { return UserAuthentication{} }
func (s SpecialUser) Grants() []AuthorizationGrant { panic("todo") }
func (s SpecialUser) History() []HistoryItem { panic("todo") }
func (s SpecialUser) GetName() string { return s.Name() }
func (s SpecialUser) Name() string { return s.Id.String() }

type RegularUser struct { types.RegularUser }
func (r RegularUser) Authentication() UserAuthentication { panic("todo") }
func (s RegularUser) History() []HistoryItem { panic("todo") }
func (r RegularUser) Grants() []AuthorizationGrant { panic("todo") }