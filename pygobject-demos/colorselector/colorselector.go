package main

import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir3/cairo"
	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

type ColorSelectorApp struct {
	window gtk.Window
	da     gtk.DrawingArea
	color  gdk.RGBA
}

func newColorSelectorApp() {
	//gdk.Color{}
	app := &ColorSelectorApp{}

	color := gdk.RGBA{P: gi.SliceAlloc0(gdk.SizeOfStructRGBA)}
	color.SetAlpha(1)
	color.SetRed(1)
	color.SetGreen(0)
	color.SetBlue(0)
	app.color = color

	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	app.window = window
	window.SetTitle("Color Selection")
	window.SetBorderWidth(8)
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	vbox := gtk.NewVBox(false, 8)
	vbox.SetBorderWidth(8)
	window.Add(vbox)

	// create color s
	frame := gtk.NewFrame("frame")
	frame.SetShadowType(gtk.ShadowTypeEtchedIn)
	vbox.PackStart(frame, true, true, 0)

	da := gtk.NewDrawingArea()
	app.da = da
	da.Connect(gtk.SigDraw, drawCb)

	da.SetSizeRequest(200, 200)
	// set the color
	da.OverrideBackgroundColor(0, app.color)
	frame.Add(da)

	alignment := gtk.NewAlignment(1.0, 0.5, 0.0, 0.0)
	button := gtk.NewButtonWithLabel("_Change the above color")
	button.SetUseUnderline(true)
	alignment.Add(button)
	vbox.PackStart(alignment, false, false, 0)

	button.Connect(gtk.SigClicked, app.changeColorCb)
	window.ShowAll()
}

func drawCb(p gi.ParamBox) {
	var s struct {
		Widget gtk.Window
		Cr     unsafe.Pointer
	}
	err := p.StoreStruct(&s)
	if err != nil {
		log.Fatal(err)
	}
	style := s.Widget.GetStyleContext()

	bgColor := gdk.RGBA{P: gi.SliceAlloc0(gdk.SizeOfStructRGBA)}
	defer gi.SliceFree(gdk.SizeOfStructRGBA, bgColor.P)

	style.GetBackgroundColor(0, bgColor)
	ctx := cairo.WrapContext(s.Cr)
	ctx.SetSourceRGBA(bgColor.Red(), bgColor.Green(), bgColor.Blue(), bgColor.Alpha())
	ctx.Paint()
}

func (app *ColorSelectorApp) changeColorCb() {
	dialog := gtk.NewColorSelectionDialog("Changing color")
	dialog.SetTransientFor(app.window)

	colorsel := gtk.WrapColorSelection(dialog.GetColorSelection().P)

	colorsel.SetPreviousRgba(app.color)
	colorsel.SetCurrentRgba(app.color)
	colorsel.SetHasPalette(true)

	response := dialog.Run()
	if response == int32(gtk.ResponseTypeOk) {
		colorsel.GetCurrentRgba(app.color)
		app.da.OverrideBackgroundColor(0, app.color)
		red := app.color.Red()
		green := app.color.Green()
		blue := app.color.Blue()
		log.Println("color", red, green, blue)

	}
	dialog.Destroy()
}

func main() {
	gtk.Init(0, 0)
	newColorSelectorApp()
	gtk.Main()
}
