package entity

import (
	context "github.com/tauraamui/metacanvas/ctx"
)

type Renderable interface {
	Render(*context.Context)
}
