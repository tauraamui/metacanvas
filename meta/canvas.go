package meta

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type Canvas struct {
	ctx   context.Context
	page  *page
	input *input.Pointer
}

func NewCanvas() *Canvas {
	c := &Canvas{
		ctx:  context.Context{Scale: 1},
		page: NewA4(),
	}
	c.input = &input.Pointer{}
	c.input.PointerEventTag = c.input
	return c
}

func (c *Canvas) Update(gtx layout.Context) {
	pointer.CursorNameOp{Name: c.updateInput(gtx)}.Add(gtx.Ops)
	c.page.Update(gtx)
}

func (c *Canvas) Render(gtx layout.Context) {
	c.ctx.Ops = gtx.Ops
	c.ctx.ApplyTransformsToOps()
	c.page.Render(c.ctx)
}

func (c *Canvas) updateInput(gtx layout.Context) pointer.CursorName {
	c.input.Update(gtx)

	if c.input.Drag {
		c.ctx.OffsetX -= c.input.DragDeltaX
		c.ctx.OffsetY -= c.input.DragDeltaY
		return pointer.CursorGrab
	}

	if c.input.Scroll {
		c.ctx.Scale -= c.input.ScrollY
		if c.ctx.Scale > 1 {
			c.ctx.Scale = 1
		} else if c.ctx.Scale < 0.01 {
			c.ctx.Scale = 0.01
		}
	}

	return pointer.CursorDefault
}
