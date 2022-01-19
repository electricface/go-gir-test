package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/electricface/go-gir/gdkpixbuf-2.0"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gi"
	"github.com/electricface/go-gir/gtk-3.0"
)

type actionEntry struct {
	name        string
	stockId     string
	label       string
	accelerator string
	tooltip     string
	isActive    bool // toggle action
	value       int  // radio action
	onActivate  func(action gtk.Action)
	//onChange    func(action gtk.RadioAction, current int32)
}

func activateAction(action gtk.Action) {
	name := action.GetName()
	log.Println("activate action", name)
	if name == "DarkTheme" {
		action := gtk.WrapToggleAction(action.P)
		value := action.GetActive()
		settings := gtk.SettingsGetDefault1()
		gv, err := g.NewValueWith(value)
		if err != nil {
			log.Fatal(err)
		}
		defer gv.Free()
		settings.SetProperty("gtk-application-prefer-dark-theme", gv)
		return
	}

	dialog := gtk.NewMessageDialog(_window, gtk.DialogFlagsDestroyWithParent, gtk.MessageTypeInfo,
		gtk.ButtonsTypeOk, "You activated action: \"%s\" of type %s",
		name, "?Type")
	dialog.Connect(gtk.SigResponse, func(args []interface{}) {
		var dialog0 gtk.Dialog
		var responseId int32
		err := gi.Store(args, &dialog0, &responseId)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("on response", responseId)
		dialog0.Destroy()
	})
	dialog.Show()
}

func onRadioActionChanged(action, current gtk.RadioAction) {
	if _messageLabel.P == nil {
		return
	}
	if action == current {
		name := current.GetName()
		value := current.GetCurrentValue()
		text := fmt.Sprintf("You activated radio action %s of type %s.\n"+
			" Current value: %d", name, "TYPE", value)
		_messageLabel.SetText(text)
		_infoBar.SetMessageType(gtk.MessageTypeEnum(value))
		_infoBar.Show()
		log.Println("set text", text)
	}
}

func updateStatusBar(buffer gtk.TextBuffer, statusBar gtk.Statusbar) {
	statusBar.Pop(0)
	count := buffer.GetCharCount()

	iter := gtk.TextIter{P: gi.Malloc0(gtk.SizeOfStructTextIter)}
	buffer.GetIterAtMark(iter, buffer.GetInsert())
	row := iter.GetLine()
	col := iter.GetLineOffset()
	msg := fmt.Sprintf("Cursor at row %d column %d - %d chars in document",
		row, col, count)
	statusBar.Push(0, msg)
	gi.Free(iter.P)
	iter.P = nil
}

func aboutCb(action gtk.Action) {
	log.Println("about cb")

	authors := []string{"John (J5) Palmieri",
		"Tomeu Vizoso",
		"and many more..."}
	documenters := []string{"David Malcolm",
		"Zack Goldberg",
		"and many more..."}
	license := `
This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Library General Public License as
published by the Free Software Foundation; either version 2 of the
License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Library General Public License for more details.

You should have received a copy of the GNU Library General Public
License along with the Gnome Library; see the file COPYING.LIB.  If not,
write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330,
Boston, MA 02111-1307, USA.
`
	filename := filepath.Join("data", "gtk-logo-rgb.gif")
	pixBuf, err := gdkpixbuf.NewPixbufFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	transparent := pixBuf.AddAlpha(true, 0xff, 0xff, 0xff)

	about := gtk.NewAboutDialog()
	about.SetTransientFor(_window)
	about.SetProgramName("Gtk+ Code Demos")
	about.SetVersion("0.1")
	about.SetCopyright("(C) 2010 The PyGI Team")
	about.SetLicense(license)
	about.SetWebsite("http://live.gnome.org/PyGI")
	about.SetComments("Program to demonstrate PyGI functions.")

	// authors
	authorsArr := gi.NewCStrArrayZTWithStrings(authors...)
	defer authorsArr.FreeAll()
	about.SetAuthors(authorsArr)

	// documenters
	documentersArr := gi.NewCStrArrayZTWithStrings(documenters...)
	defer documentersArr.FreeAll()
	about.SetDocumenters(documentersArr)

	about.SetLogo(transparent)
	about.SetTitle("About GTK+ Code Demos")
	about.Connect(gtk.SigResponse, about.Destroy)
	about.Show()
}

var actionEntries = []actionEntry{
	{name: "FileMenu", label: "_File"},
	{name: "OpenMenu", label: "_Open"},
	{name: "PreferencesMenu", label: "_Preferences"},
	{name: "ColorMenu", label: "_Color"},
	{name: "ShapeMenu", label: "_Shape"},
	{name: "HelpMenu", label: "_Help"},
	{name: "New", stockId: gtk.STOCK_NEW, label: "_New", accelerator: "<control>N",
		tooltip: "Create a new file", onActivate: activateAction},
	{name: "File1", label: "File1", tooltip: "Open first file", onActivate: activateAction},
	{name: "Save", stockId: gtk.STOCK_SAVE, label: "_Save", accelerator: "<control>S",
		tooltip: "Save current file", onActivate: activateAction},
	{name: "SaveAs", stockId: gtk.STOCK_SAVE_AS, label: "Save _As...",
		tooltip: "Save to a file", onActivate: activateAction},
	{name: "Quit", stockId: gtk.STOCK_QUIT, label: "_Quit",
		accelerator: "<control>Q", tooltip: "Quit", onActivate: activateAction},
	{name: "About", label: "_About", accelerator: "<control>A",
		tooltip: "About", onActivate: aboutCb},
	{name: "Logo", stockId: "demo-gtk-logo", tooltip: "GTK+", onActivate: activateAction},
}

func addActions(group gtk.ActionGroup, entries []actionEntry) {
	for _, entry := range entries {
		entryCopy := entry
		action := gtk.NewAction(entry.name, entry.label, entry.tooltip, entry.stockId)
		if entryCopy.onActivate != nil {
			action.Connect(gtk.SigActivate, func(args []interface{}) {
				var action0 gtk.Action
				err := gi.Store(args, &action0)
				if err != nil {
					log.Fatal(err)
				}
				entryCopy.onActivate(action0)
			})
		}
		if entry.accelerator == "" {
			group.AddAction(action)
		} else {
			group.AddActionWithAccel(action, entry.accelerator)
		}
	}
}

var toggleActionEntries = []actionEntry{
	{name: "Bold", stockId: gtk.STOCK_BOLD, label: "_Bold", accelerator: "<control>B",
		tooltip: "Bold", isActive: true, onActivate: activateAction},
	{name: "DarkTheme", label: "_Prefer Dark Theme", tooltip: "Prefer Dark Theme",
		onActivate: activateAction},
}

func addToggleActions(group gtk.ActionGroup, entries []actionEntry) {
	for _, entry := range entries {
		action := gtk.NewToggleAction(entry.name, entry.label, entry.tooltip, entry.stockId)
		action.SetActive(entry.isActive)
		entryCopy := entry
		if entryCopy.onActivate != nil {
			action.Connect(gtk.SigActivate, func(args []interface{}) {
				var action0 gtk.Action
				err := gi.Store(args, &action0)
				if err != nil {
					log.Fatal(err)
				}
				entryCopy.onActivate(action0)
			})
		}
		if entry.accelerator == "" {
			group.AddAction(action)
		} else {
			group.AddActionWithAccel(action, entry.accelerator)
		}
	}
}

const (
	colorRed int = iota
	colorGreen
	colorBlue
)

var colorActionEntries = []actionEntry{
	{name: "Red", label: "_Red", accelerator: "<control>R",
		tooltip: "Blood", value: colorRed},
	{name: "Green", label: "_Green", accelerator: "<control>G",
		tooltip: "Grass", value: colorGreen},
	{name: "Blue", label: "_Blue", accelerator: "<control>B",
		tooltip: "Sky", value: colorBlue},
}

const (
	shapeSquare int = iota
	shapeRectangle
	shapeOval
)

var shapeActionEntries = []actionEntry{
	{name: "Square", label: "_Square", accelerator: "<control>S",
		tooltip: "Square", value: shapeSquare},
	{name: "Rectangle", label: "_Rectangle", accelerator: "<control>R",
		tooltip: "Rectangle", value: shapeRectangle},
	{name: "Oval", label: "_Oval", accelerator: "<control>O",
		tooltip: "Egg", value: shapeOval},
}

func addRadioActions(group gtk.ActionGroup, entries []actionEntry,
	value int, onChange func(action, current gtk.RadioAction)) {

	var lastAction gtk.RadioAction
	for _, entry := range entries {
		action := gtk.NewRadioAction(entry.name, entry.label, entry.tooltip,
			entry.stockId, int32(entry.value))
		if onChange != nil {
			action.Connect(gtk.SigChanged, func(args []interface{}) {
				var action0 gtk.RadioAction
				var current gtk.RadioAction
				err := gi.Store(args, &action0, &current)
				if err != nil {
					log.Fatal(err)
				}

				onChange(action0, current)
			})
		}
		if value == entry.value {
			action.SetActive(true)
		}
		action.JoinGroup(lastAction)
		lastAction = action

		if entry.accelerator == "" {
			group.AddAction(action)
		} else {
			group.AddActionWithAccel(action, entry.accelerator)
		}
	}
}

func registerStockIcons() {
	// TODO
	factory := gtk.NewIconFactory()
	factory.AddDefault()

	filename := filepath.Join("data", "gtk-logo-rgb.gif")
	pixBuf, err := gdkpixbuf.NewPixbufFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	transparent := pixBuf.AddAlpha(true, 0xff, 0xff, 0xff)
	iconSet := gtk.NewIconSetFromPixbuf(transparent)

	factory.Add("demo-gtk-logo", iconSet)
}

func _quit() {
	gtk.MainQuit()
}

const uiInfo = `<ui>
    <menubar name='MenuBar'>
        <menu action='FileMenu'>
            <menuitem action='New'/>
            <menuitem action='Open'/>
            <menuitem action='Save'/>
            <menuitem action='SaveAs'/>
            <separator/>
            <menuitem action='Quit'/>
        </menu>
        <menu action='PreferencesMenu'>
            <menuitem action='DarkTheme'/>
            <menu action='ColorMenu'>
                <menuitem action='Red'/>
                <menuitem action='Green'/>
                <menuitem action='Blue'/>
            </menu>
            <menu action='ShapeMenu'>
                <menuitem action='Square'/>
                <menuitem action='Rectangle'/>
                <menuitem action='Oval'/>
            </menu>
            <menuitem action='Bold'/>
        </menu>
        <menu action='HelpMenu'>
            <menuitem action='About'/>
        </menu>
    </menubar>
    <toolbar name='ToolBar'>
        <toolitem action='Open'>
            <menu action='OpenMenu'>
                <menuitem action='File1'/>
            </menu>
        </toolitem>
        <toolitem action='Quit'/>
        <separator action='Sep1'/>
        <toolitem action='Logo'/>
    </toolbar>
</ui>
`

var _window gtk.Window
var _infoBar gtk.InfoBar
var _messageLabel gtk.Label

func main() {
	registerStockIcons()
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	_window = window
	window.SetTitle("Application Window")
	window.SetIconName("gtk-open")
	window.ConnectAfter(gtk.SigDestroy, _quit)
	table := gtk.NewTable(1, 5, false)
	window.Add(table)

	actionGroup := gtk.NewActionGroup("AppWindowActions")
	openAction := gtk.NewAction("Open", "_Open", "Open a file",
		gtk.STOCK_OPEN)
	actionGroup.AddAction(openAction)

	addActions(actionGroup, actionEntries)
	addToggleActions(actionGroup, toggleActionEntries)
	addRadioActions(actionGroup, colorActionEntries, colorRed, onRadioActionChanged)
	addRadioActions(actionGroup, shapeActionEntries, shapeSquare, onRadioActionChanged)

	merge := gtk.NewUIManager()
	merge.InsertActionGroup(actionGroup, 0)
	window.AddAccelGroup(merge.GetAccelGroup())

	_, err := merge.AddUiFromString(uiInfo, int64(len(uiInfo)))
	if err != nil {
		log.Fatal(err)
	}
	bar := merge.GetWidget("/MenuBar")
	bar.Show()
	table.Attach(bar, 0, 1, 0, 1,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		0, 0, 0)
	bar = merge.GetWidget("/ToolBar")
	bar.Show()
	table.Attach(bar, 0, 1, 1, 2,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		0, 0, 0)

	infoBar := gtk.NewInfoBar()
	_infoBar = infoBar
	infoBar.SetNoShowAll(true)
	messageLabel := gtk.NewLabel("")
	_messageLabel = messageLabel
	messageLabel.Show()

	gtk.WrapContainer(infoBar.GetContentArea().P).Add(messageLabel)
	infoBar.GetContentArea()
	infoBar.AddButton(gtk.STOCK_OK, int32(gtk.ResponseTypeOk))
	infoBar.Connect(gtk.SigResponse, func(args []interface{}) {
		var infoBar0 gtk.InfoBar
		var responseId int32
		err := gi.Store(args, &infoBar0, &responseId)
		if err != nil {
			log.Fatal(err)
		}
		infoBar0.Hide()
	})

	table.Attach(infoBar, 0, 1, 2, 3,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		0, 0, 0)
	sw := gtk.NewScrolledWindow(nil, nil)
	sw.SetShadowType(gtk.ShadowTypeIn)
	table.Attach(sw, 0, 1, 3, 4,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		0, 0)

	contents := gtk.NewTextView()
	contents.GrabFocus()
	sw.Add(contents)

	// Create status bar
	statusBar := gtk.NewStatusbar()
	table.Attach(statusBar, 0, 1, 4, 5,
		gtk.AttachOptionsExpand|gtk.AttachOptionsFill,
		0, 0, 0)

	// show text widget info in the statusbar
	buffer := contents.GetBuffer()
	buffer.Connect(gtk.SigChanged, func() {
		updateStatusBar(buffer, statusBar)
	})
	buffer.Connect(gtk.SigMarkSet, func() {
		updateStatusBar(buffer, statusBar)
	})

	updateStatusBar(buffer, statusBar)
	window.SetDefaultSize(500, 200)
	window.ShowAll()
	gtk.Main()
}
