package input

import (
	"image"

	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/op/clip"
	context "github.com/tauraamui/metacanvas/ctx"
)

type Pointer struct {
	AOE             f32.Rectangle
	PointerEventTag interface{}
	PointerID       pointer.ID
	Scroll          bool
	ScrollY         float32
	Pressed         bool
	Dragging        bool
	DragDelta       f32.Point
	Position        f32.Point
}

func (i *Pointer) Update(ctx *context.Context, eq event.Queue) {
	i.Scroll = false
	for _, evt := range eq.Events(&i.PointerEventTag) {
		x, ok := evt.(pointer.Event)
		if !ok {
			continue
		}

		switch x.Type {
		case pointer.Scroll:
			if i.PointerID != x.PointerID {
				break
			}
			if x.Scroll.Y != 0 {
				i.Scroll = true
				i.ScrollY = x.Scroll.Y * .02
			}
		case pointer.Press:
			if i.Dragging {
				break
			}
			i.Pressed = true
			i.Dragging = false
			i.PointerID = x.PointerID
			i.Position = x.Position
		case pointer.Drag:
			if i.PointerID != x.PointerID {
				break
			}
			i.Dragging = true
			i.Pressed = false
			i.DragDelta = i.Position.Sub(x.Position)
			i.Position = x.Position
		case pointer.Release:
			fallthrough
		case pointer.Cancel:
			i.Pressed = false
			i.Dragging = false
			i.Scroll = false
		}
	}

	if !i.AOE.Empty() {
		area := clip.RRect{Rect: i.AOE}.Push(ctx.Ops)
		defer area.Pop()
	}

	pointer.InputOp{
		Tag:          &i.PointerEventTag,
		ScrollBounds: image.Rect(-1, -1, 1, 1),
		Types:        pointer.Press | pointer.Drag | pointer.Release | pointer.Scroll,
	}.Add(ctx.Ops)
}
