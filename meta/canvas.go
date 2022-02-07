package meta

import (
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type Canvas struct {
	theme   *material.Theme
	page    *page
	input   *input
	offsetX float32
	offsetY float32
}

func NewCanvas() *Canvas {
	return &Canvas{
		input: &input{pointerEventTag: struct{}{}},
		theme: material.NewTheme(gofont.Collection()),
		page:  NewA4(),
	}
}

func (c *Canvas) Render(gtx layout.Context) {
	defer c.applyTransforms(gtx).Pop()
	c.page.Render(gtx.Ops)
}

func (c *Canvas) applyTransforms(gtx layout.Context) op.TransformStack {
	c.input.update(gtx)
	if c.input.drag {
		c.offsetX -= c.input.dragDeltaX
		c.offsetY -= c.input.dragDeltaY
	}
	transforms := op.Offset(f32.Pt(c.offsetX, c.offsetY)).Push(gtx.Ops)
	return transforms
}

type input struct {
	pointerEventTag struct{}
	pointerID       pointer.ID
	drag            bool
	dragDeltaX      float32
	dragDeltaY      float32
	lastPosX        float32
	lastPosY        float32
}

func (i *input) update(gtx layout.Context) {
	for _, evt := range gtx.Events(&i.pointerEventTag) {
		x, ok := evt.(pointer.Event)
		if !ok {
			continue
		}

		switch x.Type {
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
		}
	}

	pointer.InputOp{Tag: &i.pointerEventTag,
		Types: pointer.Press | pointer.Drag | pointer.Release,
		Grab:  i.drag,
	}.Add(gtx.Ops)
}
