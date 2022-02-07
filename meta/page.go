package meta

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type page struct {
	W, H int
}

func NewA4() *page {
	return &page{
		W: 800, H: 1000,
	}
}

func (p *page) Render(ops *op.Ops) {
	clip.Rect{
		Min: image.Pt(0, 0),
		Max: image.Point{p.W, p.H},
	}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
