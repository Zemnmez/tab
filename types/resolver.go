package types

import (
	"context"
	"time"

	"github.com/zemnmez/tab/user/authn/oidc"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) AuthorizationGrant() AuthorizationGrantResolver {
	return &authorizationGrantResolver{r}
}
func (r *Resolver) IDToken() IDTokenResolver {
	return &iDTokenResolver{r}
}
func (r *Resolver) Item() ItemResolver {
	return &itemResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) OIDCProvider() OIDCProviderResolver {
	return &oIDCProviderResolver{r}
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

func (r *authorizationGrantResolver) From(ctx context.Context, obj *AuthorizationGrant) (User, error) {
	panic("not implemented")
}
func (r *authorizationGrantResolver) Valid(ctx context.Context, obj *AuthorizationGrant) (*bool, error) {
	panic("not implemented")
}

type iDTokenResolver struct{ *Resolver }

func (r *iDTokenResolver) AuthorizedParty(ctx context.Context, obj *oidc.IDToken) (*string, error) {
	panic("not implemented")
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Parent(ctx context.Context, obj *Item) (*Item, error) {
	panic("not implemented")
}
func (r *itemResolver) Children(ctx context.Context, obj *Item) ([]*Item, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Ok(ctx context.Context) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Authentication(ctx context.Context) (*AuthenticationMutation, error) {
	panic("not implemented")
}
func (r *mutationResolver) Item(ctx context.Context, id ItemID, new *ItemInput) (*Item, error) {
	panic("not implemented")
}
func (r *mutationResolver) User(ctx context.Context) (*UserMutation, error) {
	panic("not implemented")
}

type oIDCProviderResolver struct{ *Resolver }

func (r *oIDCProviderResolver) ID(ctx context.Context, obj *oidc.Provider) (*string, error) {
	panic("not implemented")
}
func (r *oIDCProviderResolver) Callback(ctx context.Context, obj *oidc.Provider) (string, error) {
	panic("not implemented")
}
func (r *oIDCProviderResolver) AuthorizationEndpoint(ctx context.Context, obj *oidc.Provider) (string, error) {
	panic("not implemented")
}
func (r *oIDCProviderResolver) ClientID(ctx context.Context, obj *oidc.Provider) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ok(ctx context.Context) (*bool, error) {
	panic("not implemented")
}
func (r *queryResolver) Authentication(ctx context.Context) (*AuthenticationQuery, error) {
	panic("not implemented")
}
func (r *queryResolver) Item(ctx context.Context, id ItemID) (*Item, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context) (*UserQuery, error) {
	panic("not implemented")
}

type regularUserResolver struct{ *Resolver }

func (r *regularUserResolver) ID(ctx context.Context, obj *RegularUser) (*RegularUserID, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Created(ctx context.Context, obj *RegularUser) (*time.Time, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Authentication(ctx context.Context, obj *RegularUser) (*UserAuthentication, error) {
	panic("not implemented")
}
func (r *regularUserResolver) Grants(ctx context.Context, obj *RegularUser) ([]*AuthorizationGrant, error) {
	panic("not implemented")
}
func (r *regularUserResolver) History(ctx context.Context, obj *RegularUser) ([]*HistoryItem, error) {
	panic("not implemented")
}

type specialUserResolver struct{ *Resolver }

func (r *specialUserResolver) Name(ctx context.Context, obj *SpecialUser) (string, error) {
	panic("not implemented")
}
func (r *specialUserResolver) Authentication(ctx context.Context, obj *SpecialUser) (*UserAuthentication, error) {
	panic("not implemented")
}
func (r *specialUserResolver) Grants(ctx context.Context, obj *SpecialUser) ([]*AuthorizationGrant, error) {
	panic("not implemented")
}
func (r *specialUserResolver) History(ctx context.Context, obj *SpecialUser) ([]*HistoryItem, error) {
	panic("not implemented")
}
