package entity

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type TextBox struct {
	X, Y, W, H float32
}

func (t *TextBox) id() uint {
	return 0
}

func (t TextBox) Render(ops *op.Ops) {
	bounds := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(t.X, t.Y),
			Max: f32.Pt(t.X+t.W, t.Y+t.H),
		},
	}

	// outline
	cs := clip.Outline{Path: bounds.Path(ops)}.Op().Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, B: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cs.Pop()
}
