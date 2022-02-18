package ctx

import (
	"gioui.org/f32"
	"gioui.org/op"
)

type Context struct {
	dirty    bool
	Ops      *op.Ops
	MinScale float32
	MaxScale float32
	scale    float32
	offset   f32.Point
}

func New(scale float32) *Context {
	return &Context{scale: scale}
}

func (c *Context) IsDirty() bool {
	y := c.dirty
	c.dirty = false
	return y
}

func (c *Context) Pt2Screen(pt f32.Point) f32.Point {
	return pt.Sub(c.offset)
}

func (c *Context) Screen2Pt(pt f32.Point) f32.Point {
	return pt.Add(c.offset)
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

func (c *Context) ApplyTransformsToOps() {
	op.Offset(c.offset).Add(c.Ops)
	aff := f32.Affine2D{}.Scale(
		f32.Pt(0, 0),
		f32.Pt(c.scale, c.scale),
	)
	op.Affine(aff).Add(c.Ops)
}
