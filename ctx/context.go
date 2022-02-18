package ctx

import (
	"gioui.org/f32"
	"gioui.org/op"
)

type Context struct {
	Ops     *op.Ops
	Scale   float32
	OffsetX float32
	OffsetY float32
}

func (c *Context) Pt2Screen(pt f32.Point) f32.Point {
	return pt.Sub(f32.Pt(c.OffsetX, c.OffsetY))
}

func (c *Context) Screen2Pt(pt f32.Point) f32.Point {
	return pt.Add(f32.Pt(c.OffsetX, c.OffsetY))
}

func (c *Context) ApplyTransformsToOps() {
	op.Offset(f32.Pt(c.OffsetX, c.OffsetY)).Add(c.Ops)
	aff := f32.Affine2D{}.Scale(
		f32.Pt(0, 0),
		f32.Pt(c.Scale, c.Scale),
	)
	op.Affine(aff).Add(c.Ops)
}
