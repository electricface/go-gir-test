package main

import (
	"log"
	"unsafe"

	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("GTK Events")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	window.Connect(gtk.SigKeyPressEvent, func(args []interface{}) {
		eventP := args[1].(unsafe.Pointer)
		ev := gdk.Event{P: eventP}
		ok, keyCode := ev.GetKeycode()
		if ok {
			log.Println("key press", keyCode)
		}
	})
	window.Connect(gtk.SigMotionNotifyEvent, func(args []interface{}) {
		eventP := args[1].(unsafe.Pointer)
		ev := gdk.Event{P: eventP}
		isEv, x, y := ev.GetCoords()
		if isEv {
			log.Println(x, y)
		}
	})

	window.SetEvents(int32(gdk.EventMaskPointerMotionMask | gdk.EventMaskPointerMotionHintMask | gdk.EventMaskButtonPressMask))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
