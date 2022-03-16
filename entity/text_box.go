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
	bounds clip.RRect
	anchor anchor
	active bool
}

func NewTextBox(min, max f32.Point) *TextBox {
	return &TextBox{
		anchor: anchor{position: min},
		bounds: clip.RRect{
			Rect: f32.Rectangle{Min: min, Max: max},
		},
	}
}

// structure to contain functionality related to tracking the cursor's movement
// relative to the bounds and position of the textbox area
type anchor struct {
	position        f32.Point
	point           f32.Point
	recentlyPressed bool
}

func (a *anchor) update(ip *input.Pointer, inBounds bool) (f32.Point, bool) {
	if ip.Pressed {
		a.updatePressed(ip, inBounds)
		return f32.Point{}, false
	}

	if ip.Dragging {
		return a.updateDragged(ip), true
	}

	return f32.Point{}, false
}

func (a *anchor) updateDragged(ip *input.Pointer) f32.Point {
	delta := ip.Position.Sub(a.position).Sub(a.point)
	return delta
}

func (a *anchor) updatePressed(ip *input.Pointer, inBounds bool) {
	if !inBounds {
		a.point = f32.Point{}
		return
	}

	if !a.recentlyPressed {
		a.recentlyPressed = true
	}

	if a.recentlyPressed {
		a.recentlyPressed = false
		a.point = ip.Position.Sub(a.position)
		return
	}

	a.point = f32.Point{}
}

func (t *TextBox) id() uint {
	return 0
}

func (t *TextBox) Update(ctx *context.Context, ip *input.Pointer) bool {
	inBounds := t.withinBounds(ctx.Aff, ctx.ScreenToPt(ip.Position))
	cursor, captured := t.anchorUpdate(ip, inBounds), t.active
	pointer.CursorNameOp{Name: cursor}.Add(ctx.Ops)
	return captured
}

func (t *TextBox) Render(ctx *context.Context) {
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

func (t *TextBox) anchorUpdate(ip *input.Pointer, inBounds bool) pointer.CursorName {
	delta, moving := t.anchor.update(ip, inBounds)
	if moving {
		t.bounds.Rect.Min = t.bounds.Rect.Min.Add(delta)
		t.anchor.position = t.bounds.Rect.Min
		t.bounds.Rect.Max = t.bounds.Rect.Max.Add(delta)
		return pointer.CursorGrab
	}

	if inBounds {
		if ip.Pressed {
			t.active = true
		}
		return pointer.CursorText
	} else {
		if ip.Pressed {
			t.active = false
		}
	}

	return pointer.CursorDefault
}

func (t *TextBox) withinBounds(aff f32.Affine2D, p f32.Point) bool {
	return p.In(transformRect(t.bounds.Rect, aff))
}

func transformRect(r f32.Rectangle, aff f32.Affine2D) f32.Rectangle {
	return f32.Rectangle{Min: aff.Transform(r.Min), Max: aff.Transform(r.Max)}
}
