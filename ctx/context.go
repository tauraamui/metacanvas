package ctx

import (
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/op"
)

type Context struct {
	dirty    bool
	Ops      *op.Ops
	Events   event.Queue
	MinScale float32
	MaxScale float32
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

func (c *Context) Pt2Screen(pt f32.Point) f32.Point {
	return pt.Add(c.offset)
}

func (c *Context) Screen2Pt(pt f32.Point) f32.Point {
	return pt.Sub(c.offset)
}

func (c *Context) ScreenRect2PtRect(x, y, w, h float32) f32.Rectangle {
	x = x + c.scale
	y = y + c.scale
	return f32.Rect(x, y, x+(w*c.scale), y+(h*c.scale)).Add(c.offset)
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

func (c *Context) ApplyTransformsToOps() op.TransformStack {
	op.Offset(c.offset).Push(c.Ops)
	scale := c.scale
	aff := f32.Affine2D{}.Scale(
		c.pos,
		f32.Pt(scale, scale),
	)
	return op.Affine(aff).Push(c.Ops)
}
