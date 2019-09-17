// Package table exposes variants of types with bindings to tables
package table

import (
	"github.com/zemnmez/tab/types"
	"github.com/zemnmez/tab/storage/table"
)


type HistoryItem struct { types.HistoryItem }
func (h HistoryItem) Key() table.Key { return types.Key { tableID: types.TABLE_HISTORY_ITEM, Id: h.Id.Id } }

type SingletonUser struct { types.SingletonUser }
func (s SingletonUser) Key() table.Key { return types.Key { tableID: types.TABLE_SINGLETON_USER, Id: h.Id.Id } }

type User struct { types.User }
func (u User) Key() table.Key { return types.Key { tableID: types.TABLE_USER, Id: u.Id.Id } }

type Item struct { types.Item }
func (i Item) Key() table.Key { return types.Key { tableID: types.TABLE_ITEM, Id: i.Id.Id } }

type Provider struct { types.Provider }
func (p Provider) Key() table.Key { return types.Key { tableID: types.TABLE_PROVIDER, Id: p.Id.Id } }

type AuthzToken struct { types.AuthzToken }
func (a AuthzToken) Key() table.Key { return types.Key { tableID: types.TABLE_AUTHZ_TOKEN, Id: a.Id.Id } }

type Link struct { types.Link }
func (l Link) Table() io.Writer { return types.Key { types.TABLE_LINK } }