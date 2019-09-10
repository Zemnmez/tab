package resolver

import (
	"github.com/zemnmez/tab/types"
)

type AuthenticationMutation struct{}
type AuthenticationQuery struct{}
type UserAuthentication struct{}

func (UserAuthentication) Etc() Etc { return Etc{} }
func (AuthenticationMutation) Etc() Etc { return Etc{} }
func (AuthenticationQuery) Etc() Etc { return Etc{} }

type AuthorizationGrant struct { types.AuthorizationGrant }
func (a AuthorizationGrant) From() User { panic("todo") }
func (a AuthorizationGrant) Valid() bool { panic("todo") }