package main

import "github.com/linuxdeepin/go-gir/gtk-3.0"

const title = "Button Boxes"
const description = `
The Button Box widgets are used to arrange buttons with padding.
`

func newButtonBoxApp() {
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("Button Boxes")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)
	window.SetBorderWidth(10)

	mainVBox := gtk.NewVBox(false, 0)
	window.Add(mainVBox)

	frameHorz := gtk.NewFrame("Horizontal Button Boxes")
	mainVBox.PackStart(frameHorz, true, true, 10)

	vbox := gtk.NewVBox(false, 0)
	vbox.SetBorderWidth(10)
	frameHorz.Add(vbox)

	vbox.PackStart(createBBox(true, "Spread", 40, gtk.ButtonBoxStyleSpread),
		true, true, 0)

	vbox.PackStart(createBBox(true, "Edge", 40, gtk.ButtonBoxStyleEdge),
		true, true, 5)

	vbox.PackStart(createBBox(true, "Start", 40, gtk.ButtonBoxStyleStart),
		true, true, 5)

	vbox.PackStart(createBBox(true, "End", 40, gtk.ButtonBoxStyleEnd),
		true, true, 5)

	frameVert := gtk.NewFrame("Vertical Button Boxes")
	mainVBox.PackStart(frameVert, true, true, 10)

	hbox := gtk.NewHBox(false, 0)
	hbox.SetBorderWidth(10)
	frameVert.Add(hbox)

	hbox.PackStart(createBBox(false, "Spread", 40, gtk.ButtonBoxStyleSpread),
		true, true, 0)

	hbox.PackStart(createBBox(false, "Edge", 40, gtk.ButtonBoxStyleEdge),
		true, true, 5)

	hbox.PackStart(createBBox(false, "Start", 40, gtk.ButtonBoxStyleStart),
		true, true, 5)

	hbox.PackStart(createBBox(false, "End", 40, gtk.ButtonBoxStyleEnd),
		true, true, 5)

	window.ShowAll()
}

func bboxAssign(v *gtk.ButtonBox, src gtk.IButtonBox) {
	v.P = src.P_ButtonBox()
}

func createBBox(isHorizontal bool, title string, spacing int32, layout gtk.ButtonBoxStyleEnum) gtk.Widget {
	frame := gtk.NewFrame(title)
	var bbox gtk.ButtonBox
	if isHorizontal {
		//bbox = gtk.WrapButtonBox(gtk.NewHButtonBox().P)
		//bbox.Assign(gtk.NewHButtonBox())
		bboxAssign(&bbox, gtk.NewHButtonBox())
	} else {
		//bbox = gtk.WrapButtonBox(gtk.NewVButtonBox().P)
		bboxAssign(&bbox, gtk.NewVButtonBox())
	}
	bbox.SetBorderWidth(5)
	frame.Add(bbox)

	bbox.SetLayout(layout)
	bbox.SetSpacing(spacing)

	button := gtk.NewButtonFromStock(gtk.STOCK_OK)
	bbox.Add(button)

	button = gtk.NewButtonFromStock(gtk.STOCK_CANCEL)
	bbox.Add(button)

	button = gtk.NewButtonFromStock(gtk.STOCK_HELP)
	bbox.Add(button)

	return frame.Widget
}

func main() {
	gtk.Init(0, 0)
	newButtonBoxApp()
	gtk.Main()
}
