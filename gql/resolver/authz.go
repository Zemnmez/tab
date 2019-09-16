package resolver

type AuthorizationTokenInfo struct { types.AuthorizationToken }
type AuthorizationRefreshTokenInfo struct { type.AuthorizationToken }


type AuthorizationQuery {} 
func (AuthenticationQuery) Token(ctx context.Context, a AuthorizationToken) (inf AuthorizationTokenInfo, err error) {
	var authorizationTokenID AuthorizationTokenID
	if id, err = a.ID(); err != nil { return }

	if auth, err = Context.Get(nil, ctx).Storage.Get(id, &inf); err != nil { return }

	return

}

// TODO: think about whether i give a shit
func (a AuthorizationToken) ID() (id AuthorizationID, err error) { return AuthorizationID(a), nil }


func (OIDCMutation) Authenticate(token IDTokenInput, client AuthClientID, scopes ...Authorization) (tok AuthorizationToken, err error) {
	
}

