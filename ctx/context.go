package ctx

import (
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
)

type Context struct {
	Ctx      layout.Context
	Ops      *op.Ops
	Aff      f32.Affine2D
	Events   event.Queue
	MinScale float32
	MaxScale float32
	dirty    bool
	pos      f32.Point
	scale    float32
	offset   f32.Point
}

func DefaultCtx() *Context {
	return &Context{
		pos:      f32.Pt(0, 0),
		scale:    1,
		MinScale: 0.001,
		MaxScale: 10,
	}
}

func (c *Context) IsDirty() bool {
	y := c.dirty
	c.dirty = false
	return y
}

func (c *Context) PtToScreen(pt f32.Point) f32.Point {
	return pt.Add(c.pos).Add(c.offset)
}

func (c *Context) ScreenToPt(pt f32.Point) f32.Point {
	return pt.Add(c.pos).Sub(c.offset)
}

func (c *Context) PtRectToScreen() f32.Rectangle {
	return f32.Rectangle{}
}

func (c *Context) SetOffset(o f32.Point) {
	c.offset = o
	c.dirty = true
}

func (c *Context) AddOffset(o f32.Point) {
	c.offset = c.offset.Add(o)
	c.dirty = true
}

func (c *Context) SubOffset(o f32.Point) {
	c.offset = c.offset.Sub(o)
	c.dirty = true
}

func (c *Context) SetScale(s float32) {
	if s < c.MinScale {
		c.scale = c.MinScale
		return
	}

	if s > c.MaxScale {
		c.scale = c.MaxScale
		return
	}

	c.scale = s
	c.dirty = true
}

func (c *Context) AddScale(s float32) {
	if c.scale+s > c.MaxScale {
		c.scale = c.MaxScale
		return
	}
	c.scale += s
	c.dirty = true
}

func (c *Context) SubScale(s float32) {
	if c.scale-s < c.MinScale {
		c.scale = c.MinScale
		return
	}
	c.scale -= s
	c.dirty = true
}

func (c *Context) Scale() float32 {
	return c.scale
}

func (c *Context) ApplyTransformsToOps() op.TransformStack {
	op.Offset(c.offset).Push(c.Ops)
	scale := c.scale
	c.Aff = f32.Affine2D{}.Scale(
		c.pos,
		f32.Pt(scale, scale),
	)

	return op.Affine(c.Aff).Push(c.Ops)
}
