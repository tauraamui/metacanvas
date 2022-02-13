package meta

import (
	"image"

	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/tauraamui/metacanvas/entity"
)

type Canvas struct {
	theme   *material.Theme
	page    *page
	input   *input
	scale   float32
	offsetX float32
	offsetY float32
	ee      entity.Stack
}

func NewCanvas() *Canvas {
	return &Canvas{
		input: &input{pointerEventTag: struct{}{}},
		theme: material.NewTheme(gofont.Collection()),
		scale: 1,
		page:  NewA4(),
	}
}

func (c *Canvas) Update(gtx layout.Context) {
	pointer.CursorNameOp{Name: c.updateInput(gtx)}.Add(gtx.Ops)
}

func (c *Canvas) Render(gtx layout.Context) {
	c.applyTransforms(gtx)
	c.ee.Render(gtx.Ops)
	c.page.Render(gtx.Ops)
}

func (c *Canvas) updateInput(gtx layout.Context) pointer.CursorName {
	c.input.update(gtx)

	if c.input.drag {
		c.offsetX -= c.input.dragDeltaX
		c.offsetY -= c.input.dragDeltaY
		return pointer.CursorGrab
	}

	if c.input.scroll {
		c.scale -= c.input.scrollY
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

type input struct {
	pointerEventTag struct{}
	pointerID       pointer.ID
	scroll          bool
	scrollY         float32
	drag            bool
	dragDeltaX      float32
	dragDeltaY      float32
	lastPosX        float32
	lastPosY        float32
}

func (i *input) update(gtx layout.Context) {
	i.scroll = false
	for _, evt := range gtx.Events(&i.pointerEventTag) {
		x, ok := evt.(pointer.Event)
		if !ok {
			continue
		}

		switch x.Type {
		case pointer.Scroll:
			if i.pointerID != x.PointerID {
				break
			}
			if x.Scroll.Y != 0 {
				i.scroll = true
				i.scrollY = x.Scroll.Y * .02
			}
		case pointer.Press:
			if i.drag {
				break
			}
			i.pointerID = x.PointerID
			i.lastPosX = x.Position.X
			i.lastPosY = x.Position.Y
		case pointer.Drag:
			if i.pointerID != x.PointerID {
				break
			}
			i.drag = true
			i.dragDeltaX = i.lastPosX - x.Position.X
			i.dragDeltaY = i.lastPosY - x.Position.Y
			i.lastPosX = x.Position.X
			i.lastPosY = x.Position.Y
		case pointer.Release:
			fallthrough
		case pointer.Cancel:
			i.drag = false
			i.scroll = false
		}
	}

	pointer.InputOp{Tag: &i.pointerEventTag,
		ScrollBounds: image.Rect(-1, -1, 1, 1),
		Types:        pointer.Press | pointer.Drag | pointer.Release | pointer.Scroll,
		Grab:         i.drag,
	}.Add(gtx.Ops)
}
