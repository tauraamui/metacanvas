package meta

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/tauraamui/metacanvas/input"
)

type Canvas struct {
	page    *page
	input   *input.Pointer
	scale   float32
	offsetX float32
	offsetY float32
}

func NewCanvas() *Canvas {
	c := &Canvas{
		scale: 1,
		page:  NewA4(),
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
	c.applyTransforms(gtx)
	c.page.Render(gtx.Ops)
}

func (c *Canvas) updateInput(gtx layout.Context) pointer.CursorName {
	c.input.Update(gtx)

	if c.input.Drag {
		c.offsetX -= c.input.DragDeltaX
		c.offsetY -= c.input.DragDeltaY
		return pointer.CursorGrab
	}

	if c.input.Scroll {
		c.scale -= c.input.ScrollY
		if c.scale > 1 {
			c.scale = 1
		} else if c.scale < 0.01 {
			c.scale = 0.01
		}
	}

	return pointer.CursorDefault
}

func (c *Canvas) applyTransforms(gtx layout.Context) {
	op.Offset(f32.Pt(c.offsetX, c.offsetY)).Add(gtx.Ops)
	aff := f32.Affine2D{}.Scale(
		f32.Pt(0, 0),
		f32.Pt(c.scale, c.scale),
	)
	op.Affine(aff).Add(gtx.Ops)
}
