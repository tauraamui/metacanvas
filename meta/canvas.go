package meta

import (
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/op"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type Canvas struct {
	yOffset      int
	ctx          *context.Context
	page         *page
	input        *input.Pointer
	WindowInsets system.Insets
}

func NewCanvas() *Canvas {
	c := &Canvas{
		ctx:  context.DefaultCtx(),
		page: NewA4(),
	}
	c.input = &input.Pointer{}
	c.input.PointerEventTag = c.input
	return c
}

func (c *Canvas) SetYOffset(o int) {
	c.yOffset = o
	c.input.YOffset = float32(o)
}

func (c *Canvas) Render(ops *op.Ops, eq event.Queue) {
	c.ctx.Ops = ops
	c.updateInput(c.ctx, eq)
	st := c.ctx.ApplyTransformsToOps()
	defer st.Pop()
	c.page.Render(c.ctx)
}

func (c *Canvas) updateInput(ctx *context.Context, eq event.Queue) {
	c.input.Update(c.ctx, eq)

	if captured := c.page.Update(ctx, c.input); captured {
		return
	}

	if c.input.Dragging {
		c.ctx.AddOffset(c.input.DragDelta.Mul(2.002))
		pointer.CursorNameOp{Name: pointer.CursorGrab}.Add(ctx.Ops)
		return
	}

	if c.input.Scroll {
		c.ctx.SubScale(c.input.ScrollY)
	}
}
