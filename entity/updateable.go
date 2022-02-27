package entity

import (
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type Updateable interface {
	Update(*context.Context, *input.Pointer) bool
}
