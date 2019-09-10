package resolver

import (
	"github.com/zemnmez/tab/types"
)

type Item struct { types.Item }
func (i Item) Parent() *Item { panic("todo") }
func (i Item) Children() []*Item { panic("todo") } 