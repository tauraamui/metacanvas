package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"github.com/tacusci/logging/v2"
	"github.com/tauraamui/metacanvas/meta"
)

func main() {
	go func() {
		w := app.NewWindow()
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
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			ops.Reset()
			canvas.Render(&ops, e.Queue)
			e.Frame(&ops)
		}
	}
}
