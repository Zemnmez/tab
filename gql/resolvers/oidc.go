package resolvers

import "github.com/zemnmez/tab/types"
import "github.com/zemnmez/tab/gql/models"

func todo() { panic("unimplemented") }

type OIDCMutation struct{}
type OIDCQuery struct{}
type OIDCProviderQuery struct{}

func (OIDCMutation) Provider(id types.OIDCProviderID, provider OIDCProviderInput) types.OIDCProvider {
	todo()
}

func (OIDCMutation) Authenticate(token IDTokenInput) User           { todo() }
func (OIDCQuery) IsValid(token IDTokenInput) bool                   { todo() }
func (OIDCProviderQuery) All() []types.OIDCProvider                 { todo() }
func (OIDCProviderQuery) ByID(id types.OIDCProviderID) types.OIDCProvider { todo() }

func (AuthenticationMutation) OIDC() OIDCMutation { return OIDCMutation{} }
func (AuthenticationQuery) OIDCQuery() OIDCQuery            { return OIDCQuery{} }
func (OIDCQuery) Provider() OIDCProviderQuery     { return OIDCProviderQuery{} }
func (UserAuthentication) OIDC() []types.IDToken  { todo() }
