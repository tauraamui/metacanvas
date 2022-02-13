package entity

import (
	"gioui.org/op"
)

type Renderable interface {
	Render(*op.Ops)
}
