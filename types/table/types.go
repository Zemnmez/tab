// Package table exposes variants of types with bindings to tables
package table

import (
	"github.com/zemnmez/tab/types"
)

type HistoryItem { types.HistoryItem }
func (HistoryItem) Table() io.Writer { return types.TABLE_HISTORY_ITEM }
func (h HistoryItem) ID() io.Writer { return h.HistoryItem.Id }

type SingletonUser { types.SingletonUser }
func (SingletonUser) Table() io.Writer { return types.TABLE_SINGLETON_USER }
func (s SingletonUser) ID() io.Writer { return s.SingletonUser.Id }

type User { types.User }
func (User) Table() io.Writer { return types.TABLE_USER }
func (u User) ID() io.Writer { return u.User.Id }

type Item { types.Item }
func (Item) Table() io.Writer { return types.TABLE_ITEM }
func (i Item) ID() io.Writer { return i.Item.Id }

type Provider { types.Provider }
func (Provider) Table() io.Writer { return types.TABLE_OIDC_PROVIDER }
func (p Provider) ID() io.Writer { return p.Provider.Id }

type AuthzToken { types.AuthzToken }
func (AuthzToken) Table() io.Writer { return types.TABLE_AUTHZ_TOKEN }
func (a AuthzToken) ID() io.Writer { return p.AuthzToken.Id }

type Link { types.Link }
func (Link) Table() io.Writer { return types.TABLE_LINK }
func (l Link) ID() io.Writer { return p.Link.Id }