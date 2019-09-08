package graphql

import (
	"context"
	"time"

	"github.com/zemnmez/tab/types"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) AuthorizationGrant() AuthorizationGrantResolver {
	return &authorizationGrantResolver{r}
}
func (r *Resolver) Item() ItemResolver {
	return &itemResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) RegularUser() RegularUserResolver {
	return &regularUserResolver{r}
}
func (r *Resolver) SpecialUser() SpecialUserResolver {
	return &specialUserResolver{r}
}

type authorizationGrantResolver struct{ *Resolver }

func (r *authorizationGrantResolver) From(ctx context.Context, obj *types.AuthorizationGrant) (User, error) {
	panic("not implemented")
}
func (r *authorizationGrantResolver) Valid(ctx context.Context, obj *types.AuthorizationGrant) (*bool, error) {
	panic("not implemented")
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Parent(ctx context.Context, obj *types.Item) (*types.Item, error) {
	panic("not implemented")
}
func (r *itemResolver) Children(ctx context.Context, obj *types.Item) ([]*types.Item, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Ok(ctx context.Context) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Authentication(ctx context.Context) (*AuthenticationMutation, error) {
	panic("not implemented")
}
func (r *mutationResolver) Item(ctx context.Context, id types.ItemID, new *ItemInput) (*types.Item, error) {
	panic("not implemented")
}
func (r *mutationResolver) User(ctx context.Context) (*UserMutation, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ok(ctx context.Context) (*bool, error) {
	panic("not implemented")
}
func (r *queryResolver) Authentication(ctx context.Context) (*AuthenticationQuery, error) {
	panic("not implemented")
}
func (r *queryResolver) Item(ctx context.Context, id types.ItemID) (*types.Item, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context) (*UserQuery, error) {
	panic("not implemented")
}

type regularUserResolver struct{ *Resolver }

func (r *regularUserResolver) ID(ctx context.Context, obj *types.RegularUser) (*types.RegularUserID, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Created(ctx context.Context, obj *types.RegularUser) (*time.Time, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Authentication(ctx context.Context, obj *types.RegularUser) (*UserAuthentication, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Grants(ctx context.Context, obj *types.RegularUser) ([]*types.AuthorizationGrant, error) {
	panic("not implemented")
}
func (r *regularUserResolver) History(ctx context.Context, obj *types.RegularUser) ([]*HistoryItem, error) {
	panic("not implemented")
}

type specialUserResolver struct{ *Resolver }

func (r *specialUserResolver) ID(ctx context.Context, obj *types.SpecialUser) (SpecialUserID, error) {
	panic("not implemented")
}
func (r *specialUserResolver) Name(ctx context.Context, obj *types.SpecialUser) (string, error) {
	panic("not implemented")
}
func (r *specialUserResolver) Authentication(ctx context.Context, obj *types.SpecialUser) (*UserAuthentication, error) {
	panic("not implemented")
}
func (r *specialUserResolver) Grants(ctx context.Context, obj *types.SpecialUser) ([]*types.AuthorizationGrant, error) {
	panic("not implemented")
}
func (r *specialUserResolver) History(ctx context.Context, obj *types.SpecialUser) ([]*HistoryItem, error) {
	panic("not implemented")
}
