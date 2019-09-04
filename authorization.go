package tab

type AuthorizationGrant struct {
	From UserID
	Of   []Authorization
}

func (a AuthorizationGrant) Valid() bool { panic("unimplemented") }

func (UserQuery) WhoCan(do []Authorization) ([]User, error) {
	panic("unimplemented")
}

func (UserMutation) GrantRegular(to RegularUserID, abilities []Authorization) (RegularUser, error) {
	panic("unimplemented")
}

func (UserMutation) GrantSpecial(to SpecialUserID, abilities []Authorization) (SpecialUser, error) {
	panic("unimplemented")
}
