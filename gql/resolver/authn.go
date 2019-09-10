package resolver

type AuthenticationMutation struct{}
type AuthenticationQuery struct{}
type UserAuthentication struct{}

func (UserAuthentication) Etc() Etc { return Etc{} }
func (AuthenticationMutation) Etc() Etc { return Etc{} }
func (AuthenticationQuery) Etc() Etc { return Etc{} }
