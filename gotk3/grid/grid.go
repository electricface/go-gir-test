package main

import (
	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.SetTitle("Grid Example")
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	grid := gtk.NewGrid()

	grid.SetOrientation(gtk.OrientationVertical)
	lab := gtk.NewLabel("Just a label")
	btn := gtk.NewButtonWithLabel("Button with label")
	entry := gtk.NewEntry()
	spnBtn := gtk.NewSpinButtonWithRange(0, 1.0, 0.001)
	nb := gtk.NewNotebook()

	grid.Add(btn)
	grid.Add(lab)
	grid.Add(entry)
	grid.Add(spnBtn)

	grid.Attach(nb, 1, 1, 1, 2)
	nb.SetHexpand(true)
	nb.SetVexpand(true)

	nbChild := gtk.NewLabel("Notebook content")
	nbTab := gtk.NewLabel("Tab label")
	nb.AppendPage(nbChild, nbTab)

	win.Add(grid)
	win.ShowAll()

	gtk.Main()
}
