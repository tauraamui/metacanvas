package entity

import (
	"fmt"
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	context "github.com/tauraamui/metacanvas/ctx"
	"github.com/tauraamui/metacanvas/input"
)

type TextBox struct {
	X, Y, W, H float32
	bounds     clip.RRect
	pressed    bool
}

func (t *TextBox) id() uint {
	return 0
}

func (t *TextBox) Update(ctx *context.Context, ip *input.Pointer) bool {
	return pointer.CursorDefault != t.updateInput(ctx, ip)
}

func swapShade(t bool) color.NRGBA {
	if t {
		return color.NRGBA{R: 0xff, A: 0xff}
	}
	return color.NRGBA{R: 0x80, B: 0x80, A: 0xFF}
}

func (t *TextBox) Render(ctx *context.Context) {
	t.bounds = clip.RRect{
		Rect: f32.Rectangle{
			Min: f32.Pt(t.X, t.Y),
			Max: f32.Pt(t.X+t.W, t.Y+t.H),
		},
	}

	// outline
	cs := clip.Outline{Path: t.bounds.Path(ctx.Ops)}.Op().Push(ctx.Ops)
	paint.ColorOp{Color: swapShade(t.pressed)}.Add(ctx.Ops)
	paint.PaintOp{}.Add(ctx.Ops)
	cs.Pop()
}

func (t *TextBox) updateInput(ctx *context.Context, ip *input.Pointer) pointer.CursorName {
	fmt.Printf("IP: %s\n", ctx.ScreenToPt(ip.Position).String())
	fmt.Printf("TBX POS+SIZE: X/Y: %s, XW/YH: %s\n", f32.Pt(t.X, t.Y).String(), f32.Pt(t.X+t.W, t.Y+t.H))
	if ip.Pressed && t.withinBounds(ctx.ScreenToPt(ip.Position)) {
		t.pressed = true
		return pointer.CursorGrab
	}
	t.pressed = false
	return pointer.CursorDefault
}

func (t *TextBox) withinBounds(p f32.Point) bool {
	return p.In(t.bounds.Rect)
}
