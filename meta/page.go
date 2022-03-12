package meta

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/entity"
	"github.com/tauraamui/metacanvas/input"
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
			entity.NewTextBox(f32.Pt(10, 10), f32.Pt(400, 40)),
		},
	}
}

func (p *page) Update(ctx *context.Context, ip *input.Pointer) bool {
	return p.ee.Update(ctx, ip)
}

func (p *page) Render(ctx *context.Context) {
	bounds := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(p.X, p.Y),
			Max: f32.Pt(p.X+p.W, p.Y+p.H),
		},
	}

	if p.clipToPage {
		bc := bounds.Push(ctx.Ops)
		defer bc.Pop()
	}

	// entities stack
	p.ee.Render(ctx)

	// page outline
	cl := clip.Stroke{Path: bounds.Path(ctx.Ops), Width: 1.2}.Op().Push(ctx.Ops)
	paint.ColorOp{Color: color.NRGBA{R: 0xFF, A: 0xFF}}.Add(ctx.Ops)
	paint.PaintOp{}.Add(ctx.Ops)
	cl.Pop()
}
