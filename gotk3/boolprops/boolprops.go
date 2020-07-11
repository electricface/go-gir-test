package main

import (
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

// Setup the Window.
func setupWindow() gtk.Window {
	w := gtk.NewWindow(gtk.WindowTypeToplevel)
	w.Connect(gtk.SigDestroy, gtk.MainQuit)
	w.SetDefaultSize(500, 300)
	w.SetPosition(gtk.WindowPositionCenter)
	w.SetTitle("TextView properties example")

	return w
}

// Setup the TextView, put it in a ScrolledWindow, and add both to box.
func setupTextView(box gtk.Box) gtk.TextView {
	sw := gtk.NewScrolledWindow(nil, nil)
	tv := gtk.NewTextView()
	sw.Add(tv)
	box.PackStart(sw, true, true, 0)
	return tv
}

type BoolProperty struct {
	Name string
	Get  func() bool
	Set  func(bool)
}

func setupPropertyCheckboxes(outer gtk.Box, props []*BoolProperty) {
	box := gtk.NewBox(gtk.OrientationHorizontal, 0)
	for _, prop := range props {
		chk := gtk.NewCheckButtonWithLabel(prop.Name)
		// initialize the checkbox with the property's current value
		chk.SetActive(prop.Get())
		p := prop // w/o this all the checkboxes will toggle the last property in props
		chk.Connect(gtk.SigToggled, func() {
			p.Set(chk.GetActive())
		})
		box.PackStart(chk, true, true, 0)
	}
	outer.PackStart(box, false, false, 0)
}

func main() {
	gtk.Init(0, 0)

	win := setupWindow()
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	win.Add(box)

	tv := setupTextView(box)

	props := []*BoolProperty{
		{"cursor visible", tv.GetCursorVisible, tv.SetCursorVisible},
		{"editable", tv.GetEditable, tv.SetEditable},
		{"overwrite", tv.GetOverwrite, tv.SetOverwrite},
		{"accepts tab", tv.GetAcceptsTab, tv.SetAcceptsTab},
	}

	setupPropertyCheckboxes(box, props)

	win.ShowAll()

	gtk.Main()
}
