package entity

import (
	"gioui.org/layout"
	"gioui.org/op"
)

type Stack []Entity

func (s Stack) Update(gtx layout.Context) {
	for _, e := range s {
		e.Update(gtx)
	}
}

func (s Stack) Render(ops *op.Ops) {
	for _, e := range s {
		e.Render(ops)
	}
}
