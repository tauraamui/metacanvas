package main

import (
	"image"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/tacusci/logging/v2"
	"github.com/tauraamui/metacanvas/meta"
)

const WIN_WIDTH, WIN_HEIGHT = 900, 700

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("metacanvas"),
			app.Size(unit.Dp(WIN_WIDTH), unit.Dp(WIN_HEIGHT)),
		)
		err := run(w)
		if err != nil {
			logging.Fatal(err.Error())
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	canvas := meta.NewCanvas()
	var decorationHeight int
	var initalWindowSize image.Point
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case app.ConfigEvent:
			if initalWindowSize.X == 0 || initalWindowSize.Y == 0 {
				initalWindowSize = e.Config.Size
			}
		case system.FrameEvent:
			ops.Reset()

			if decorationHeight == 0 && initalWindowSize.X > 0 && initalWindowSize.Y > 0 {
				decorationHeight = initalWindowSize.Y - e.Size.Y
				canvas.SetYOffset(decorationHeight)
			}

			canvas.Render(&ops, e.Queue)
			e.Frame(&ops)
		}
	}
}
