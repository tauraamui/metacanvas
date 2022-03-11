package entity

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type TextBox struct {
	Min, Max f32.Point
	bounds   clip.RRect
	active   bool
}

func (t *TextBox) id() uint {
	return 0
}

func (t *TextBox) Update(ctx *context.Context, ip *input.Pointer) bool {
	return pointer.CursorDefault != t.updateInput(ctx, ip)
}

func (t *TextBox) Render(ctx *context.Context) {
	t.bounds = clip.RRect{
		Rect: f32.Rectangle{
			Min: t.Min,
			Max: t.Max,
		},
	}

	if t.active {
		t.renderOutline(ctx)
	}
}

func (t *TextBox) renderOutline(ctx *context.Context) {
	cl := clip.Stroke{Path: t.bounds.Path(ctx.Ops), Width: 1.3}.Op().Push(ctx.Ops)
	paint.ColorOp{Color: color.NRGBA{A: 0xFF}}.Add(ctx.Ops)
	paint.PaintOp{}.Add(ctx.Ops)
	cl.Pop()
}

func (t *TextBox) updateInput(ctx *context.Context, ip *input.Pointer) pointer.CursorName {
	inBounds := t.withinBounds(ctx.Aff, ctx.ScreenToPt(ip.Position))
	if inBounds {
		if ip.Pressed {
			t.active = true
			return pointer.CursorDefault
		}
		return pointer.CursorText
	}

	if ip.Pressed {
		t.active = false
		return pointer.CursorDefault
	}

	return pointer.CursorDefault
}

func (t *TextBox) withinBounds(aff f32.Affine2D, p f32.Point) bool {
	return p.In(transformRect(t.bounds.Rect, aff))
}

func transformRect(r f32.Rectangle, aff f32.Affine2D) f32.Rectangle {
	return f32.Rectangle{Min: aff.Transform(r.Min), Max: aff.Transform(r.Max)}
}
