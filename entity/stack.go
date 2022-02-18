package entity

import (
	"gioui.org/layout"
	context "github.com/tauraamui/metacanvas/ctx"
)

type Stack []Entity

func (s Stack) Update(gtx layout.Context) {
	for _, e := range s {
		e.Update(gtx)
	}
}

func (s Stack) Render(ctx *context.Context) {
	for _, e := range s {
		e.Render(ctx)
	}
}
