// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"

	"github.com/zemnmez/tab"
	"github.com/zemnmez/tab/user"
	"github.com/zemnmez/tab/user/authn"
	"github.com/zemnmez/tab/user/authn/oidc"
	"github.com/zemnmez/tab/user/authz"
)

type UserMutator interface {
	IsUserMutator()
}

type AnonymousUser struct {
	ID   SpecialUserID `json:"ID"`
	Name string        `json:"Name"`
}

type AuthenticationMutation struct {
	Etc  *string       `json:"etc"`
	Oidc *OIDCMutation `json:"OIDC"`
}

type AuthenticationQuery struct {
	Etc  *string    `json:"etc"`
	Oidc *OIDCQuery `json:"OIDC"`
}

type DefinedItemInput struct {
	Name     string       `json:"Name"`
	Location string       `json:"Location"`
	Parent   *ItemInput   `json:"Parent"`
	Children []*ItemInput `json:"Children"`
}

type HistoryItem struct {
	Action      string    `json:"Action"`
	By          user.User `json:"By"`
	RequestData string    `json:"RequestData"`
	IPAddress   string    `json:"IPAddress"`
}

type ItemInput struct {
	ByID          *tab.ItemID       `json:"ById"`
	WithStructure *DefinedItemInput `json:"WithStructure"`
}

type OIDCMutation struct {
	Authenticate user.User      `json:"Authenticate"`
	Provider     *oidc.Provider `json:"Provider"`
}

type OIDCProviderInput struct {
	Name string `json:"Name"`
}

type OIDCProviderQuery struct {
	All  []*oidc.Provider `json:"All"`
	ByID *oidc.Provider   `json:"ByID"`
}

type OIDCQuery struct {
	Provider *OIDCProviderQuery `json:"Provider"`
	IsValid  *bool              `json:"IsValid"`
}

type RootUser struct {
	ID   SpecialUserID `json:"ID"`
	Name string        `json:"Name"`
}

type Self struct {
	Name           string                `json:"Name"`
	Authentication *authn.Authentication `json:"Authentication"`
	Grants         []*authz.Grant        `json:"Grants"`
	Authorizatons  []Authorization       `json:"Authorizatons"`
	// Grant a user some ability the current user has
	Grant user.User `json:"Grant"`
	// Grant a special user some ability the current user has
	GrantSpecial *user.Special  `json:"GrantSpecial"`
	History      []*HistoryItem `json:"History"`
}

func (Self) IsUser() {}

type UserInput struct {
	Name string `json:"Name"`
}

// Authorization is a list of all the possible permissions
// a User can have.
type Authorization string

const (
	AuthorizationViewUsers             Authorization = "VIEW_USERS"
	AuthorizationModifyValidAuth       Authorization = "MODIFY_VALID_AUTH"
	AuthorizationAddItems              Authorization = "ADD_ITEMS"
	AuthorizationModifyItems           Authorization = "MODIFY_ITEMS"
	AuthorizationModifyOtherUsers      Authorization = "MODIFY_OTHER_USERS"
	AuthorizationModifySpecialUsers    Authorization = "MODIFY_SPECIAL_USERS"
	AuthorizationModifySelf            Authorization = "MODIFY_SELF"
	AuthorizationViewOtherUsersHistory Authorization = "VIEW_OTHER_USERS_HISTORY"
	AuthorizationViewOwnHistory        Authorization = "VIEW_OWN_HISTORY"
)

var AllAuthorization = []Authorization{
	AuthorizationViewUsers,
	AuthorizationModifyValidAuth,
	AuthorizationAddItems,
	AuthorizationModifyItems,
	AuthorizationModifyOtherUsers,
	AuthorizationModifySpecialUsers,
	AuthorizationModifySelf,
	AuthorizationViewOtherUsersHistory,
	AuthorizationViewOwnHistory,
}

func (e Authorization) IsValid() bool {
	switch e {
	case AuthorizationViewUsers, AuthorizationModifyValidAuth, AuthorizationAddItems, AuthorizationModifyItems, AuthorizationModifyOtherUsers, AuthorizationModifySpecialUsers, AuthorizationModifySelf, AuthorizationViewOtherUsersHistory, AuthorizationViewOwnHistory:
		return true
	}
	return false
}

func (e Authorization) String() string {
	return string(e)
}

func (e *Authorization) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Authorization(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Authorization", str)
	}
	return nil
}

func (e Authorization) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SpecialUserID string

const (
	SpecialUserIDRoot      SpecialUserID = "ROOT"
	SpecialUserIDAnonymous SpecialUserID = "ANONYMOUS"
)

var AllSpecialUserID = []SpecialUserID{
	SpecialUserIDRoot,
	SpecialUserIDAnonymous,
}

func (e SpecialUserID) IsValid() bool {
	switch e {
	case SpecialUserIDRoot, SpecialUserIDAnonymous:
		return true
	}
	return false
}

func (e SpecialUserID) String() string {
	return string(e)
}

func (e *SpecialUserID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SpecialUserID(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SpecialUserID", str)
	}
	return nil
}

func (e SpecialUserID) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
