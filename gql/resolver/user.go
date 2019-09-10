package resolver

import (
	"context"

	"github.com/zemnmez/tab/types"
)

type UserQuery struct{}
type UserMutation struct{}

func (UserQuery) Self(ctx context.Context) Self                             { panic("todo") }
func (UserQuery) Special(id types.SpecialUserID) types.SpecialUser          { panic("todo") }
func (UserQuery) Regular(id types.RegularUserID) types.RegularUser          { panic("todo") }
func (UserQuery) WhoCan(do []types.Authorization) (users []User, err error) { panic("todo") }

func (UserMutation) Self() UserMutator { panic("todo") }
func (UserMutation) Special() UserMutator { panic("todo") }
func (UserMutation) Regular() UserMutator { panic("todo") }

func (Mutation) User() UserMutation { return UserMutation{} }

type UserMutator interface {
	Modify(with UserInput) User
}

type User interface {
	GetName() string
}

func (s Self) GetName() string { return s.Name }
