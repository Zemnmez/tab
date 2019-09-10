package resolver

import (
	"github.com/zemnmez/tab/types"
)

type HistoryItem struct { types.HistoryItem }
func (h HistoryItem) By() User { panic("todo") }