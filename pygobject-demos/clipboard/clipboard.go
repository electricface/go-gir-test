package main

import (
	"log"

	"github.com/electricface/go-gir/gdk-3.0"
	"github.com/electricface/go-gir/gdkpixbuf-2.0"
	"github.com/electricface/go-gir/gi"
	"github.com/electricface/go-gir/gtk-3.0"
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
	//_dataMap["entry1"] = entry
	button.Connect(gtk.SigClicked, copyButtonClicked, entry)

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
	button.Connect(gtk.SigClicked, pasteButtonClicked, entry)

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
	ebox.Connect(gtk.SigDragBegin, dragBegin, image)
	ebox.Connect(gtk.SigDragDataGet, dragDataGet, image)

	// accept drops on ebox
	ebox.DragDestSet(gtk.DestDefaultsAll, nil, 0, gdk.DragActionCopy)
	ebox.DragDestAddImageTargets()
	ebox.Connect(gtk.SigDragDataReceived, dragDataReceived, image)

	// context menu on ebox
	ebox.Connect(gtk.SigButtonPressEvent, buttonPress, image)

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
	ebox.Connect(gtk.SigDragBegin, dragBegin, image)
	ebox.Connect(gtk.SigDragDataGet, dragDataGet, image)

	// accept drops on ebox
	ebox.DragDestSet(gtk.DestDefaultsAll, nil, 0, gdk.DragActionCopy)
	ebox.DragDestAddImageTargets()
	ebox.Connect(gtk.SigDragDataReceived, dragDataReceived, image)

	// context menu on ebox
	ebox.Connect(gtk.SigButtonPressEvent, buttonPress, image)

	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	clipboard.SetCanStore(nil, 0)

	window.ShowAll()
}

func buttonPress(p gi.ParamBox) interface{} {
	var s struct {
		Widget gtk.Widget
		Event  gdk.Event
	}
	err := p.StoreStruct(&s)
	if err != nil {
		log.Fatal(err)
	}
	_, button := s.Event.GetButton()
	if button != 3 {
		return false
	}

	image := p.UserData.(gtk.Image)
	menu := gtk.NewMenu()

	item := gtk.NewImageMenuItemFromStock(gtk.STOCK_COPY, nil)
	item.Connect(gtk.SigActivate, copyImage, image)
	item.Show()
	menu.Append(item)

	item = gtk.NewImageMenuItemFromStock(gtk.STOCK_PASTE, nil)
	item.Connect(gtk.SigActivate, pasteImage, image)
	item.Show()
	menu.Append(item)

	time1 := s.Event.GetTime()
	menu.Popup(nil, nil, nil, button, time1)
	return false
}

func copyImage(p gi.ParamBox) {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	image := p.UserData.(gtk.Image)
	pixbuf := getImagePixBuf(image)
	clipboard.SetImage(pixbuf)
}

func pasteImage(p gi.ParamBox) {
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := gtk.ClipboardGet1(atom)
	pixBuf := clipboard.WaitForImage()

	if pixBuf.P != nil {
		image := p.UserData.(gtk.Image)
		image.SetFromPixbuf(pixBuf)
	}
}

func dragBegin(p gi.ParamBox) {
	var widget gtk.Widget
	var ctx gdk.DragContext
	err := p.Store(&widget, &ctx)
	if err != nil {
		log.Fatal(err)
	}
	image := p.UserData.(gtk.Image)
	pixBuf := getImagePixBuf(image)
	gtk.DragSetIconPixbuf(ctx, pixBuf, -2, -2)
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

func dragDataGet(p gi.ParamBox) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	err := p.StoreStruct(&s)
	if err != nil {
		log.Fatal(err)
	}
	image := p.UserData.(gtk.Image)
	pixBuf := getImagePixBuf(image)
	s.Data.SetPixbuf(pixBuf)
}

func dragDataReceived(p gi.ParamBox) {
	var s struct {
		Widget gtk.Widget
		Ctx    gdk.DragContext
		X      int32
		Y      int32
		Data   gtk.SelectionData
		Info   uint32
		Time   uint32
	}
	err := p.StoreStruct(&s)
	if err != nil {
		log.Fatal(err)
	}
	if s.Data.GetLength() > 0 {
		pixBuf := s.Data.GetPixbuf()
		image := p.UserData.(gtk.Image)
		image.SetFromPixbuf(pixBuf)
	}
}

func copyButtonClicked(p gi.ParamBox) {
	entry := p.UserData.(gtk.Entry)
	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := entry.GetClipboard(atom)

	// set the clipboard's text
	clipboard.SetText(entry.GetText(), -1)
}

func pasteButtonClicked(p gi.ParamBox) {
	entry := p.UserData.(gtk.Entry)

	// get the default clipboard
	atom := gdk.AtomIntern("CLIPBOARD", true)
	clipboard := entry.GetClipboard(atom)

	// set the clipboard's text
	clipboard.RequestText(func(clipboard gtk.Clipboard, text string) {
		if text != "" {
			entry.SetText(text)
		}
	})
}

func main() {
	gtk.Init(0, 0)
	newClipboardApp()
	gtk.Main()
}
