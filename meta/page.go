package meta

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/tauraamui/metacanvas/entity"
)

type page struct {
	X, Y, W, H float32
	clipToPage bool
	ee         entity.Stack
}

func NewA4() *page {
	return &page{
		W: 800, H: 1000,
		ee: entity.Stack{
			&entity.TextBox{X: 10, Y: 10, W: 250, H: 20},
		},
	}
}

func (p *page) Update(gtx layout.Context) {
	p.ee.Update(gtx)
}

func (p *page) Render(ops *op.Ops) {
	bounds := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(0, 0),
			Max: f32.Pt(p.W, p.H),
		},
	}

	if p.clipToPage {
		bc := bounds.Push(ops)
		defer bc.Pop()
	}

	// entities stack
	p.ee.Render(ops)

	// page outline
	cl := clip.Stroke{Path: bounds.Path(ops), Width: 1.2}.Op().Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0xFF, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()
}
