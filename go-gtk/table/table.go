package main

import (
	"fmt"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("GTK Table")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.PolicyTypeAutomatic, gtk.PolicyTypeAutomatic)

	table := gtk.NewTable(5, 5, false)
	for y := uint(0); y < 50; y++ {
		for x := uint(0); x < 50; x++ {
			table.Attach(gtk.NewButtonWithLabel(fmt.Sprintf("%02d:%02d", x, y)), uint32(x), uint32(x+1), uint32(y), uint32(y+1), gtk.AttachOptionsFill, gtk.AttachOptionsFill, 5, 5)
		}
	}
	swin.AddWithViewport(table)

	window.Add(swin)
	window.SetDefaultSize(200, 200)
	window.ShowAll()

	gtk.Main()
}
