package entity

import (
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type Stack []Entity

func (s Stack) Update(ctx *context.Context, ip *input.Pointer) bool {
	for _, e := range s {
		if captured := e.Update(ctx, ip); captured {
			return captured
		}
	}
	return false
}

func (s Stack) Render(ctx *context.Context) {
	for _, e := range s {
		e.Render(ctx)
	}
}
