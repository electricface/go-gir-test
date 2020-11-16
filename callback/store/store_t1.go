package main

import (
	"fmt"
	"log"

	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gi"

	"github.com/davecgh/go-spew/spew"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	w := gtk.NewWindow(gtk.WindowTypeToplevel)
	w.Show()
	w.SetSizeRequest(100, 100)
	w.Connect(gtk.SigDestroy, gtk.MainQuit)
	w.Connect(gtk.SigButtonPressEvent, func(args []interface{}) interface{} {
		log.Println("button press")
		spew.Dump(args)
		var arg0 gtk.Window
		var ev gdk.Event
		err := gi.Store(args, &arg0, &ev)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("arg0: %#v, arg1: %#v\n", arg0, ev)
		_, btn := ev.GetButton()
		_, x, y := ev.GetCoords()
		log.Println(x, y)
		log.Println(btn)

		var args0 struct {
			Window gtk.Window
			Ev     gdk.Event
		}
		err = gi.StoreStruct(args, &args0)
		if err != nil {
			log.Fatal(err)
		}
		ev = args0.Ev
		_, btn = ev.GetButton()
		_, x, y = ev.GetCoords()
		args0.Window.SetTitle(fmt.Sprint(btn, x, y))
		log.Println("- ", x, y)
		log.Println("- ", btn)
		return true
	})
	gtk.Main()
}
