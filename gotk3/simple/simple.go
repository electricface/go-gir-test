package main

import (
	"github.com/electricface/go-gir/gtk-3.0"
	"log"
)

func main() {
	gtk.Init(0, 0)
	log.Println("hello")

	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.SetTitle("simple example")
	win.Connect(gtk.SigDestroy, gtk.MainQuit)

	l := gtk.NewLabel("Hello, gotk3")
	win.Add(l)
	win.SetDefaultSize(800, 800)
	win.ShowAll()

	gtk.Main()
}
