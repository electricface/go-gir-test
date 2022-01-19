package main

import (
	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("We love Expander")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	expander := gtk.NewExpander("dan the ...")
	expander.Add(gtk.NewLabel("404 contents not found"))
	vbox.PackStart(expander, false, false, 0)

	window.Add(vbox)
	window.ShowAll()

	gtk.Main()
}
