package meta

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type page struct {
	W, H float32
}

func NewA4() *page {
	return &page{
		W: 800, H: 1000,
	}
}

func (p *page) Render(ops *op.Ops) {
	r := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(0, 0),
			Max: f32.Pt(p.W, p.H),
		},
	}
	clip.Stroke{Path: r.Path(ops), Width: 1.25}.Op().Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
