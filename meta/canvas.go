package meta

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/tauraamui/metacanvas/input"
)

type Context struct {
	ops     *op.Ops
	scale   float32
	offsetX float32
	offsetY float32
}

func (c *Context) ApplyTransformsToOps() {
	op.Offset(f32.Pt(c.offsetX, c.offsetY)).Add(c.ops)
	aff := f32.Affine2D{}.Scale(
		f32.Pt(0, 0),
		f32.Pt(c.scale, c.scale),
	)
	op.Affine(aff).Add(c.ops)
}

type Canvas struct {
	ctx   Context
	page  *page
	input *input.Pointer
}

func NewCanvas() *Canvas {
	c := &Canvas{
		ctx:  Context{scale: 1},
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
	c.ctx.ops = gtx.Ops
	c.ctx.ApplyTransformsToOps()
	c.page.Render(c.ctx)
}

func (c *Canvas) updateInput(gtx layout.Context) pointer.CursorName {
	c.input.Update(gtx)

	if c.input.Drag {
		c.ctx.offsetX -= c.input.DragDeltaX
		c.ctx.offsetY -= c.input.DragDeltaY
		return pointer.CursorGrab
	}

	if c.input.Scroll {
		c.ctx.scale -= c.input.ScrollY
		if c.ctx.scale > 1 {
			c.ctx.scale = 1
		} else if c.ctx.scale < 0.01 {
			c.ctx.scale = 0.01
		}
	}

	return pointer.CursorDefault
}
