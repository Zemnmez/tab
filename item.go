package tab

import (
	"github.com/zemnmez/tab/storage"
)

// ItemID uniquely identifies an item
type ItemID struct {
	ID
}

// Item represents the item concept within Tab, which is the location of
// an inventoried item or container
type Item struct {
	ID
	Name        string
	Location    string
	ParentID    ItemID
	ChildrenIDs []ItemID
}

// Parent returns the parent item from ParentID
func (i Item) Parent() (parent *Item, err error) { err = storage.Get(i.ParentID).Into(&parent); return }

// Children returns the children of this item from ChildrenIDs
func (i Item) Children() (children []Item, err error) {
	err = storage.Get(i.ChildrenIDs...).Into(&children)
	return
}

func (Query) Item(id ItemID) (found *Item, err error)                { err = storage.Get(id).Into(&found); return }
func (Mutation) Item(id ItemID, new ItemInput) (new Item, err error) { panic("unimplemented") }
