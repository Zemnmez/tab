package resolver

import "github.com/zemnmez/tab/types"

type OIDCMutation struct{}
type OIDCQuery struct{}
type OIDCProviderQuery struct{}

func (OIDCMutation) Provider(id types.OIDCProviderID, provider OIDCProviderInput) types.OIDCProvider {
	panic("todo")
}

func (OIDCMutation) Authenticate(token IDTokenInput) User           { panic("todo") }
func (OIDCQuery) IsValid(token IDTokenInput) bool                   { panic("todo") }
func (OIDCProviderQuery) All() []types.OIDCProvider                 { panic("todo") }
func (OIDCProviderQuery) ByID(id types.OIDCProviderID) types.OIDCProvider { panic("todo") }

func (AuthenticationMutation) OIDC() OIDCMutation { return OIDCMutation{} }
func (AuthenticationQuery) OIDCQuery() OIDCQuery            { return OIDCQuery{} }
func (OIDCQuery) Provider() OIDCProviderQuery     { return OIDCProviderQuery{} }
func (UserAuthentication) OIDC() []types.IDToken  { panic("todo") }
