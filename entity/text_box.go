package entity

import (
	"fmt"
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type TextBox struct {
	X, Y, W, H float32
	input      *input.Pointer
}

func (t *TextBox) id() uint {
	return 0
}

func (t *TextBox) Update(gtx layout.Context) {
	t.updateInput(gtx)
}

func (t *TextBox) Render(ctx context.Context) {
	bounds := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(t.X, t.Y),
			Max: f32.Pt(t.X+t.W, t.Y+t.H),
		},
	}

	// outline
	cs := clip.Outline{Path: bounds.Path(ctx.Ops)}.Op().Push(ctx.Ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, B: 0x80, A: 0xFF}}.Add(ctx.Ops)
	paint.PaintOp{}.Add(ctx.Ops)
	cs.Pop()
}

func (t *TextBox) updateInput(gtx layout.Context) pointer.CursorName {
	if t.input == nil {
		t.input = &input.Pointer{
			AOE: f32.Rect(t.X, t.Y, t.X+t.W, t.Y+t.H),
		}
		t.input.PointerEventTag = t.input
	}
	t.input.Update(gtx)

	fmt.Printf("PRESSED: %f, %f\n", t.input.LastPosX, t.input.LastPosY)

	if t.input.Drag {
		t.X -= t.input.DragDeltaX
		t.Y -= t.input.DragDeltaY
		t.input.AOE = f32.Rect(t.X, t.Y, t.X+t.W, t.Y+t.H)
		return pointer.CursorGrab
	}

	return pointer.CursorDefault
}
