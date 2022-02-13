package entity

import "gioui.org/op"

type Stack []Entity

func (s Stack) Render(ops *op.Ops) {
	for _, e := range s {
		e.Render(ops)
	}
}
