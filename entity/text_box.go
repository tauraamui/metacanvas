package entity

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
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

func (t *TextBox) Update(ctx *context.Context, eq event.Queue) {
	t.updateInput(ctx, eq)
}

func swapShade(t bool) color.NRGBA {
	if t {
		return color.NRGBA{R: 0xff, A: 0xff}
	}
	return color.NRGBA{R: 0x80, B: 0x80, A: 0xFF}
}

func (t *TextBox) Render(ctx *context.Context) {
	bounds := clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(t.X, t.Y),
			Max: f32.Pt(t.X+t.W, t.Y+t.H),
		},
	}

	// outline
	cs := clip.Outline{Path: bounds.Path(ctx.Ops)}.Op().Push(ctx.Ops)
	paint.ColorOp{Color: swapShade(t.input.Pressed)}.Add(ctx.Ops)
	paint.PaintOp{}.Add(ctx.Ops)
	cs.Pop()
}

func (t *TextBox) updateInput(ctx *context.Context, eq event.Queue) pointer.CursorName {
	if t.input == nil || ctx.IsDirty() {
		t.input = &input.Pointer{
			AOE: ctx.ScreenRect2PtRect(t.X, t.Y, t.W, t.H),
		}
		t.input.PointerEventTag = t.input
	}

	t.input.Update(ctx, eq)

	return pointer.CursorDefault
}
