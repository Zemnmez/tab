// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/zemnmez/tab"
	"github.com/zemnmez/tab/users"
)

type UserMutator interface {
	IsUserMutator()
}

type AnonymousUser struct {
	ID             SpecialUserID   `json:"ID"`
	Name           string          `json:"Name"`
	Authorizations []Authorization `json:"Authorizations"`
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
	Action      string     `json:"Action"`
	By          users.User `json:"By"`
	RequestData string     `json:"RequestData"`
	IPAddress   string     `json:"IPAddress"`
}

type IDToken struct {
	Issuer                              string    `json:"Issuer"`
	Subject                             string    `json:"Subject"`
	Audience                            string    `json:"Audience"`
	Expiration                          time.Time `json:"Expiration"`
	Issued                              time.Time `json:"Issued"`
	Nonce                               string    `json:"Nonce"`
	AuthenticationContextClassReference *int      `json:"AuthenticationContextClassReference"`
	AuthenticationMethodsReference      []string  `json:"AuthenticationMethodsReference"`
	AuthorizedParty                     *string   `json:"AuthorizedParty"`
}

type IDTokenInput struct {
	Issuer                              string    `json:"Issuer"`
	Subject                             string    `json:"Subject"`
	Audience                            string    `json:"Audience"`
	Expiration                          time.Time `json:"Expiration"`
	Issued                              time.Time `json:"Issued"`
	Nonce                               string    `json:"Nonce"`
	AuthenticationContextClassReference *int      `json:"AuthenticationContextClassReference"`
	AuthenticationMethodsReference      []string  `json:"AuthenticationMethodsReference"`
	AuthorizedParty                     *string   `json:"AuthorizedParty"`
}

type ItemInput struct {
	ByID          *tab.ItemID       `json:"ById"`
	WithStructure *DefinedItemInput `json:"WithStructure"`
}

type OIDCMutation struct {
	Authenticate users.User    `json:"Authenticate"`
	Provider     *OIDCProvider `json:"Provider"`
}

type OIDCProvider struct {
	ID                    *string `json:"ID"`
	Name                  string  `json:"Name"`
	Callback              string  `json:"Callback"`
	AuthorizationEndpoint string  `json:"AuthorizationEndpoint"`
	ClientID              string  `json:"ClientID"`
}

type OIDCProviderInput struct {
	Name string `json:"Name"`
}

type OIDCProviderQuery struct {
	All  []*OIDCProvider `json:"All"`
	ByID *OIDCProvider   `json:"ByID"`
}

type OIDCQuery struct {
	Provider *OIDCProviderQuery `json:"Provider"`
	IsValid  *bool              `json:"IsValid"`
}

type RootUser struct {
	ID             SpecialUserID   `json:"ID"`
	Name           string          `json:"Name"`
	Authorizations []Authorization `json:"Authorizations"`
}

type Self struct {
	Name           string                    `json:"Name"`
	Authentication *UserAuthentication       `json:"Authentication"`
	Grants         []*tab.AuthorizationGrant `json:"Grants"`
	Authorizatons  []Authorization           `json:"Authorizatons"`
	History        []*HistoryItem            `json:"History"`
}

func (Self) IsUser() {}

type UserAuthentication struct {
	Etc  *string    `json:"etc"`
	Oidc []*IDToken `json:"OIDC"`
}

type UserInput struct {
	Name string `json:"Name"`
}

type UserMutation struct {
	Self         UserMutator    `json:"Self"`
	Special      UserMutator    `json:"Special"`
	Regular      UserMutator    `json:"Regular"`
	GrantRegular *users.Regular `json:"GrantRegular"`
	GrantSpecial *users.Special `json:"GrantSpecial"`
}

type UserQuery struct {
	Self    *Self          `json:"Self"`
	Special *users.Special `json:"Special"`
	Regular *users.Regular `json:"Regular"`
	WhoCan  []users.User   `json:"WhoCan"`
}

type Authorization string

const (
	AuthorizationViewUsers             Authorization = "VIEW_USERS"
	AuthorizationModifyValidAuth       Authorization = "MODIFY_VALID_AUTH"
	AuthorizationViewOtherUsersHistory Authorization = "VIEW_OTHER_USERS_HISTORY"
	AuthorizationViewOwnHistory        Authorization = "VIEW_OWN_HISTORY"
	AuthorizationAddItems              Authorization = "ADD_ITEMS"
	AuthorizationModifyItems           Authorization = "MODIFY_ITEMS"
	AuthorizationModifyOtherUsers      Authorization = "MODIFY_OTHER_USERS"
	AuthorizationModifySpecialUsers    Authorization = "MODIFY_SPECIAL_USERS"
	AuthorizationModifySelf            Authorization = "MODIFY_SELF"
)

var AllAuthorization = []Authorization{
	AuthorizationViewUsers,
	AuthorizationModifyValidAuth,
	AuthorizationViewOtherUsersHistory,
	AuthorizationViewOwnHistory,
	AuthorizationAddItems,
	AuthorizationModifyItems,
	AuthorizationModifyOtherUsers,
	AuthorizationModifySpecialUsers,
	AuthorizationModifySelf,
}

func (e Authorization) IsValid() bool {
	switch e {
	case AuthorizationViewUsers, AuthorizationModifyValidAuth, AuthorizationViewOtherUsersHistory, AuthorizationViewOwnHistory, AuthorizationAddItems, AuthorizationModifyItems, AuthorizationModifyOtherUsers, AuthorizationModifySpecialUsers, AuthorizationModifySelf:
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

type Error string

const (
	ErrorNotFound Error = "NOT_FOUND"
)

var AllError = []Error{
	ErrorNotFound,
}

func (e Error) IsValid() bool {
	switch e {
	case ErrorNotFound:
		return true
	}
	return false
}

func (e Error) String() string {
	return string(e)
}

func (e *Error) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Error(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Error", str)
	}
	return nil
}

func (e Error) MarshalGQL(w io.Writer) {
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
