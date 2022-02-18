package entity

import (
	"gioui.org/io/event"
	context "github.com/tauraamui/metacanvas/ctx"
)

type Updateable interface {
	Update(*context.Context, event.Queue)
}
