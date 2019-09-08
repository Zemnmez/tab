package graphql


type UserQuery struct {}
func (u UserQuery) Self(ctx context.Context) (user.User, error) {
	panic("unimplemented")
}

type UserMutation struct { }
func (u UserQuery) Self(ctx context.Context) (user.User, error) {
	panic("unimplemented")
}