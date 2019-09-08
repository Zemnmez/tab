package resolvers

import (
	"github.com/zemnmez/tab/types"
)

type UserQuery struct {}
type UserMutation struct{}


func (UserQuery) Self() Self { panic("todo") }
func (UserQuery) Special(id types.SpecialUserID) types.SpecialUser { panic("todo") }
func (UserQuery) Regular(id types.RegularUserID) types.RegularUser { panic("todo") }

func (UserMutation) Self() UserMutator { panic("todo") }

func (Mutation) User() UserMutation { return UserMutation{} }

type UserMutator interface {
	Modify(with UserInput) User
}

type User interface {
	GetName() string
}

func (s Self) GetName() string { return s.Name }