package main

import (
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func createArrowButton(at gtk.ArrowTypeEnum, st gtk.ShadowTypeEnum) gtk.Button {
	b := gtk.NewButton()
	a := gtk.NewArrow(at, st)

	b.Add(a)

	b.Show()
	a.Show()

	return b
}

func main() {
	gtk.Init(0,0)

	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("Arrow Buttons")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	box := gtk.NewHBox(false, 0)
	box.Show()
	window.Add(box)

	up := createArrowButton(gtk.ArrowTypeUp, gtk.ShadowTypeIn)
	down := createArrowButton(gtk.ArrowTypeDown, gtk.ShadowTypeOut)
	left := createArrowButton(gtk.ArrowTypeLeft, gtk.ShadowTypeEtchedIn)
	right := createArrowButton(gtk.ArrowTypeRight, gtk.ShadowTypeEtchedOut)

	box.PackStart(up, false, false, 3)
	box.PackStart(down, false, false, 3)
	box.PackStart(left, false, false, 3)
	box.PackStart(right, false, false, 3)

	up.Connect(gtk.SigClicked,func() { println("↑") })
	down.Connect(gtk.SigClicked,func() { println("↓") })
	left.Connect(gtk.SigClicked, func() { println("←") })
	right.Connect(gtk.SigClicked,func() { println("→") })

	window.Show()
	gtk.Main()
}
