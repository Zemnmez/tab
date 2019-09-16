package resolver

import (
	"github.com/zemnmez/tab/types"
)

type LinkID struct{ types.LinkID }

type Link struct{ types.Link }


func (s SelfMutation) Unlink(ctx context.Context, issuer string) (ok bool, err error) {
	return ConnectionsMutation{}.Unlink(
		From: UserQuery{}.Self(ctx).ID,
		To: issuer,
	)
}

func (s SelfMutation) Link(ctx context.Context, to IDToken) (New Link, err error) {
	panic("todo")
}


type ConnectionsQuery struct{}

func (ConnectionsQuery) Links(For UserID) (links []Link, err error) {
	panic("todo")
}

type ConnectionsMutation struct{}

func (ConnectionsMutation) Link(from UserID, to IDToken) (New Link, err error) {
	panic("todo")
}
func (ConnectionsMutation) Unlink(from UserID, to IDToken) (ok bool, err error) {
	panic("todo")
}


