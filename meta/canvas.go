package meta

import (
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/op"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type Canvas struct {
	ctx   *context.Context
	page  *page
	input *input.Pointer
}

func NewCanvas() *Canvas {
	c := &Canvas{
		ctx:  &context.Context{},
		page: NewA4(),
	}
	c.input = &input.Pointer{}
	c.input.PointerEventTag = c.input
	return c
}

func (c *Canvas) Render(ops *op.Ops, eq event.Queue) {
	c.ctx.Ops = ops
	pointer.CursorNameOp{Name: c.updateInput(c.ctx, eq)}.Add(ops)

	c.page.Update(c.ctx, eq)

	c.ctx.ApplyTransformsToOps()

	c.page.Render(c.ctx)
}

func (c *Canvas) updateInput(ctx *context.Context, eq event.Queue) pointer.CursorName {
	c.input.Update(ctx, eq)

	if c.input.Drag {
		c.ctx.SubOffset(f32.Pt(c.input.DragDeltaX, c.input.DragDeltaY))
		return pointer.CursorGrab
	}

	if c.input.Scroll {
		c.ctx.SubScale(c.input.ScrollY)
	}

	return pointer.CursorDefault
}
