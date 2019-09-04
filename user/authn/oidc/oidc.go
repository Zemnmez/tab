package oidc

type IDToken struct {
	Issuer string
	Subject string
	Audience string
	Expiration time.Time
	Issued time.Time
	Nonce string
	AuthenticationContextClassReference int
	AuthenticationMethodsReference []string
	AuthorizedParty []string
}

type ProviderID uuid.Uuid

type Provider struct {
	ID ProviderID
	Name string
}