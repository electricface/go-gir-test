package main

import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir3/cairo"
	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

func main() {
	gtk.Init(0, 0)

	// gui boilerplate
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	da := gtk.NewDrawingArea()
	win.Add(da)
	win.SetTitle("Arrow keys")
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	win.ShowAll()

	// Data
	unitSize := 20.0
	x := 0.0
	y := 0.0
	keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() { y-- },
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() { y++ },
	}

	// Event handlers
	da.Connect(gtk.SigDraw, func(args []interface{}) {
		crp := args[1].(unsafe.Pointer)
		//da gtk.DrawingArea, cr *cairo.Context
		cr := cairo.WrapContext(crp)
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Fill()
	})
	win.Connect(gtk.SigKeyPressEvent, func(args []interface{}) {
		//win *gtk.Window, ev *gdk.Event
		evp := args[1].(unsafe.Pointer)
		ev := gdk.Event{P: evp}
		keyEvent := ev
		_, kv := keyEvent.GetKeyval()
		if move, found := keyMap[uint(kv)]; found {
			log.Println("kv:", kv)
			move()
			win.QueueDraw()
		}
	})

	gtk.Main()
}
