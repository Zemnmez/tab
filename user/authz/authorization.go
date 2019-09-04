package authz

type Grant struct {
	From UserID
	Of   []Authorization
}

func (a Grant) Valid() bool { panic("unimplemented") }