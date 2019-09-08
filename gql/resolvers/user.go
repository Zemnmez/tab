package resolvers

import (
	"github.com/zemnmez/tab/types"
	"github.com/zemnmez/tab/gql/models"
)

type UserQuery struct {}
type UserMutation struct{}


func (UserQuery) Self() Self { todo() }
func (UserQuery) Special(id types.SpecialUserID) types.SpecialUser { todo() }
func (UserQuery) Regular(id types.RegularUserID) types.RegularUser { todo() }

func (UserMutation) Self() UserMutator

func (Mutation) User() UserMutation { return UserMutation{} }

type UserMutator interface {
	Modify(with UserInput) models.User
}