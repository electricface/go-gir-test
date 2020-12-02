package main

import (
	"log"

	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gdkpixbuf-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func newClipboardApp() {
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("Clipboard demo")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	vbox := gtk.NewVBox(false, 0)
	vbox.SetBorderWidth(8)
	window.Add(vbox)

	label := gtk.NewLabel("\"Copy\" will copy the text\nin the entry to the clipboard")
	vbox.PackStart(label, false, false, 0)

	hbox := gtk.NewHBox(false, 4)
	hbox.SetBorderWidth(8)
	vbox.PackStart(hbox, false, false, 0)

	// create first entry
	entry := gtk.NewEntry()
	hbox.PackStart(entry, true, true, 0)

	// create button
	button := gtk.NewButtonFromStock(gtk.STOCK_COPY)
	hbox.PackStart(button, false, false, 0)
	_dataMap["entry1"] = entry
	button.Connect(gtk.SigClicked, copyButtonClicked)

	label = gtk.NewLabel("\"Paste\" will paste the text from the clipboard to the entry")
	vbox.PackStart(label, false, false, 0)

	hbox = gtk.NewHBox(false, 4)
	hbox.SetBorderWidth(8)
	vbox.PackStart(hbox, false, false, 0)

	// create secondary entry
	entry = gtk.NewEntry()
	hbox.PackStart(entry, true, true, 0)
	// create button
	button = gtk.NewButtonFromStock(gtk.STOCK_PASTE)
	hbox.PackStart(button, false, false, 0)
	_dataMap["entry2"] = entry
	button.Connect(gtk.SigClicked, pasteButtonClicked)

	label = gtk.NewLabel("Images can be transferred via the clipboard, too")
	vbox.PackStart(label, false, false, 0)

	hbox = gtk.NewHBox(false, 4)
	hbox.SetBorderWidth(8)
	vbox.PackStart(hbox, false, false, 0)

	// create the first image
	image := gtk.NewImageFromStock(gtk.STOCK_DIALOG_WARNING, int32(gtk.IconSizeButton))
	ebox := gtk.NewEventBox()
	ebox.Add(image)
	hbox.Add(ebox)

	// make ebox a drag source
	ebox.DragSourceSet(gdk.ModifierTypeButton1Mask, nil, 0,
		gdk.DragActionCopy)
	ebox.DragSourceAddImageTargets()
	_dataMap["image1"] = image
	ebox.Connect(gtk.SigDragBegin, dragBegin)
	ebox.Connect(gtk.SigDragDataGet, dragDataGet)

	// accept drops on ebox
	ebox.DragDestSet(gtk.DestDefaultsAll, nil, 0, gdk.DragActionCopy)
	ebox.DragDestAddImageTargets()
	ebox.Connect(gtk.SigDragDataReceived, dragDataReceived)

	// context menu on ebox
	ebox.Connect(gtk.SigButtonPressEvent, buttonPress)

	// ----------------------------------
	// image 2

	// create the second image
	image = gtk.NewImageFromStock(gtk.STOCK_STOP, int32(gtk.IconSizeButton))
	ebox = gtk.NewEventBox()
	ebox.Add(image)
	hbox.Add(ebox)

	// make ebox a drag source
	ebox.DragSourceSet(gdk.ModifierTypeButton1Mask, nil, 0,
		gdk.DragActionCopy)
	ebox.DragSourceAddImageTargets()
	_dataMap["image2"] = image
	ebox.Connect(gtk.SigDragBegin, dragBegin2)
	ebox.Connect(gtk.SigDragDataGet, dragDataGet2)

	// accept drops on ebox
	ebox.DragDestSet(gtk.DestDefaultsAll, nil, 0, gdk.DragActionCopy)
	ebox.DragDestAddImageTargets()
	ebox.Connect(gtk.SigDragDataReceived, dragDataReceived2)

	// context menu on ebox
	ebox.Connect(gtk.SigButtonPressEvent, buttonPress2)

	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	clipboard.SetCanStore(nil, 0)

	window.ShowAll()
}

func buttonPress(args []interface{}) interface{} {
	var s struct {
		Widget gtk.Widget
		Event  gdk.Event
	}
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	_, button := s.Event.GetButton()
	if button != 3 {
		return false
	}
	menu := gtk.NewMenu()

	item := gtk.NewImageMenuItemFromStock(gtk.STOCK_COPY, nil)
	item.Connect(gtk.SigActivate, copyImage)
	item.Show()
	menu.Append(item)

	item = gtk.NewImageMenuItemFromStock(gtk.STOCK_PASTE, nil)
	item.Connect(gtk.SigActivate, pasteImage)
	item.Show()
	menu.Append(item)

	time1 := s.Event.GetTime()
	menu.Popup(nil, nil, nil, button, time1)
	return false
}

func buttonPress2(args []interface{}) interface{} {
	var s struct {
		Widget gtk.Widget
		Event  gdk.Event
	}
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	_, button := s.Event.GetButton()
	if button != 3 {
		return false
	}
	menu := gtk.NewMenu()

	item := gtk.NewImageMenuItemFromStock(gtk.STOCK_COPY, nil)
	item.Connect(gtk.SigActivate, copyImage2)
	item.Show()
	menu.Append(item)

	item = gtk.NewImageMenuItemFromStock(gtk.STOCK_PASTE, nil)
	item.Connect(gtk.SigActivate, pasteImage2)
	item.Show()
	menu.Append(item)

	time1 := s.Event.GetTime()
	menu.Popup(nil, nil, nil, button, time1)
	return false
}

func copyImage() {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	image := _dataMap["image1"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	clipboard.SetImage(pixbuf)
}

func pasteImage() {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	pixbuf := clipboard.WaitForImage()

	if pixbuf.P != nil {
		image := _dataMap["image1"].(gtk.Image)
		image.SetFromPixbuf(pixbuf)
	}
}

func copyImage2() {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	image := _dataMap["image2"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	clipboard.SetImage(pixbuf)
}

func pasteImage2() {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	pixbuf := clipboard.WaitForImage()

	if pixbuf.P != nil {
		image := _dataMap["image2"].(gtk.Image)
		image.SetFromPixbuf(pixbuf)
	}
}

func dragBegin(args []interface{}) {
	var widget gtk.Widget
	var ctx gdk.DragContext
	err := gi.Store(args, &widget, &ctx)
	if err != nil {
		log.Fatal(err)
	}
	image := _dataMap["image1"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	gtk.DragSetIconPixbuf(ctx, pixbuf, -2, -2)
}

func dragBegin2(args []interface{}) {
	var widget gtk.Widget
	var ctx gdk.DragContext
	err := gi.Store(args, &widget, &ctx)
	if err != nil {
		log.Fatal(err)
	}
	image := _dataMap["image2"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	gtk.DragSetIconPixbuf(ctx, pixbuf, -2, -2)
}

func getImagePixBuf(image gtk.Image) gdkpixbuf.Pixbuf {
	storageType := image.GetStorageType()
	if storageType == gtk.ImageTypePixbuf {
		return image.GetPixbuf()
	} else if storageType == gtk.ImageTypeStock {
		stockId, size := image.GetStock()
		return image.RenderIcon(stockId, size, "")
	}
	return gdkpixbuf.Pixbuf{}
}

func dragDataGet(args []interface{}) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	image := _dataMap["image1"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	s.Data.SetPixbuf(pixbuf)
}

func dragDataGet2(args []interface{}) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	image := _dataMap["image2"].(gtk.Image)
	pixbuf := getImagePixBuf(image)
	s.Data.SetPixbuf(pixbuf)
}

func dragDataReceived(args []interface{}) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		X      int32
		Y      int32
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	// var s gtk.WidgetSigArgsDragDataReceived
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	if s.Data.GetLength() > 0 {
		pixbuf := s.Data.GetPixbuf()
		image := _dataMap["image1"].(gtk.Image)
		image.SetFromPixbuf(pixbuf)
	}
}

func dragDataReceived2(args []interface{}) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		X      int32
		Y      int32
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	// var s gtk.WidgetSigArgsDragDataReceived
	err := gi.StoreStruct(args, &s)
	if err != nil {
		log.Fatal(err)
	}
	if s.Data.GetLength() > 0 {
		pixbuf := s.Data.GetPixbuf()
		image := _dataMap["image2"].(gtk.Image)
		image.SetFromPixbuf(pixbuf)
	}
}

var _dataMap = make(map[string]interface{})

func copyButtonClicked() {
	entry := _dataMap["entry1"].(gtk.Entry)
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := entry.GetClipboard(atom)

	// set the clipboard's text
	clipboard.SetText(entry.GetText(), -1)
}

func pasteButtonClicked() {
	entry := _dataMap["entry2"].(gtk.Entry)

	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := entry.GetClipboard(atom)

	// set the clipboard's text
	clipboard.RequestText(pasteReceived)
}

func pasteReceived(clipboard gtk.Clipboard, text string) {
	entry := _dataMap["entry2"].(gtk.Entry)

	if text != "" {
		entry.SetText(text)
	}
}

func main() {
	gtk.Init(0, 0)
	newClipboardApp()
	gtk.Main()
}
