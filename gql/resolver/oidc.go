package resolver

import (
	"errors"
	"time"

	"github.com/zemnmez/tab/types"
)

// b64url jwt
type IDToken string

// Validates the signature of this IDTokenInput
func (IDTokenInput) Token() IDToken {
	panic(" todo ")
}

type IDTokenInfo struct {
	Issuer     string
	Subject    string
	Audience   string
	Expiration time.Time
	Issued     time.Time
}

// TODO: maybe make the errors nice again

func (i IDTokenInfo) Authenticate(issuer string, audience string) (user string, err error) {
	switch {
	case i.Issued.After(time.Now()):
		err = errors.New("token isn't valid yet")

	case i.Expiration.Before(time.Now()):
		err = errors.New("token has expired")

	case i.Issuer != issuer:
		err = errors.New("issuer incorrect")

	case i.Audience != audience:
		err = errors.New("audience incorrect")
	}

	user = i.Subject

	return
}

type OIDCMutation struct{}
type OIDCQuery struct{}
type OIDCProviderQuery struct{}

func (OIDCMutation) Provider(id types.OIDCProviderID, provider OIDCProviderInput) types.OIDCProvider {
	panic("todo")
}

func (OIDCMutation) Authenticate(token IDTokenInput) User                 { panic("todo") }
func (OIDCQuery) IsValid(token IDTokenInput) bool                         { panic("todo") }
func (OIDCProviderQuery) All() []types.OIDCProvider                       { panic("todo") }
func (OIDCProviderQuery) ByID(id types.OIDCProviderID) types.OIDCProvider { panic("todo") }

func (AuthenticationMutation) OIDC() OIDCMutation { return OIDCMutation{} }
func (AuthenticationQuery) OIDC() OIDCQuery       { return OIDCQuery{} }
func (OIDCQuery) Provider() OIDCProviderQuery     { return OIDCProviderQuery{} }
func (UserAuthentication) OIDC() []types.IDToken  { panic("todo") }
