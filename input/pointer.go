package input

import (
	"image"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
)

type Pointer struct {
	AOE             f32.Rectangle
	PointerEventTag interface{}
	PointerID       pointer.ID
	Scroll          bool
	ScrollY         float32
	Drag            bool
	DragDeltaX      float32
	DragDeltaY      float32
	LastPosX        float32
	LastPosY        float32
}

func (i *Pointer) Update(gtx layout.Context) {
	i.Scroll = false
	for _, evt := range gtx.Events(&i.PointerEventTag) {
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
			if i.Drag {
				break
			}
			i.PointerID = x.PointerID
			i.LastPosX = x.Position.X
			i.LastPosY = x.Position.Y
		case pointer.Drag:
			if i.PointerID != x.PointerID {
				break
			}
			i.Drag = true
			i.DragDeltaX = i.LastPosX - x.Position.X
			i.DragDeltaY = i.LastPosY - x.Position.Y
			i.LastPosX = x.Position.X
			i.LastPosY = x.Position.Y
		case pointer.Release:
			fallthrough
		case pointer.Cancel:
			i.Drag = false
			i.Scroll = false
		}
	}

	if !i.AOE.Empty() {
		area := clip.RRect{Rect: i.AOE}.Push(gtx.Ops)
		defer area.Pop()
	}

	pointer.InputOp{
		Tag:          &i.PointerEventTag,
		ScrollBounds: image.Rect(-1, -1, 1, 1),
		Types:        pointer.Press | pointer.Drag | pointer.Release | pointer.Scroll,
	}.Add(gtx.Ops)
}
