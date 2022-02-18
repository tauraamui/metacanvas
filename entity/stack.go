package entity

import (
	"gioui.org/io/event"
	context "github.com/tauraamui/metacanvas/ctx"
)

type Stack []Entity

func (s Stack) Update(ctx *context.Context, eq event.Queue) {
	for _, e := range s {
		e.Update(ctx, eq)
	}
}

func (s Stack) Render(ctx *context.Context) {
	for _, e := range s {
		e.Render(ctx)
	}
}
