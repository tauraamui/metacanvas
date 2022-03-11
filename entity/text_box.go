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
	cursor, captured := t.updateInput(ctx, ip)
	pointer.CursorNameOp{Name: cursor}.Add(ctx.Ops)
	return captured
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

func (t *TextBox) updateInput(ctx *context.Context, ip *input.Pointer) (pointer.CursorName, bool) {
	inBounds := t.withinBounds(ctx.Aff, ctx.ScreenToPt(ip.Position))
	captured := false
	if inBounds {
		captured = true
		if ip.Pressed {
			t.active = true
			return pointer.CursorDefault, captured
		}

		if ip.Dragging {
			t.active = true
			t.Min = t.Min.Sub(ip.DragDelta)
			t.Max = t.Max.Sub(ip.DragDelta)
			return pointer.CursorGrab, captured
		}
		return pointer.CursorText, captured
	}

	if ip.Pressed {
		t.active = false
		return pointer.CursorDefault, captured
	}

	return pointer.CursorDefault, captured
}

func (t *TextBox) withinBounds(aff f32.Affine2D, p f32.Point) bool {
	return p.In(transformRect(t.bounds.Rect, aff))
}

func transformRect(r f32.Rectangle, aff f32.Affine2D) f32.Rectangle {
	return f32.Rectangle{Min: aff.Transform(r.Min), Max: aff.Transform(r.Max)}
}
