package main

import (
	"log"

	"github.com/linuxdeepin/go-gir/g-2.0"

	"github.com/electricface/go-gir3/gi-lite"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
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
		dialog0.Destroy()
	})
	dialog.Show()
}

func aboutCb(action gtk.Action) {
	log.Println("about cb")
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
	value int, onChange func()) {

	for _, entry := range entries {
		action := gtk.NewRadioAction(entry.name, entry.label, entry.tooltip,
			entry.stockId, int32(entry.value))
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
		if value == entry.value {
			action.SetActive(true)
			//} else {
			//	action.SetActive(false)
		}
		if entry.accelerator == "" {
			group.AddAction(action)
		} else {
			group.AddActionWithAccel(action, entry.accelerator)
		}
	}
}

func registerStockIcons() {
	// TODO
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

	// add actions action_entries
	addActions(actionGroup, actionEntries)

	// add toggle actions
	addToggleActions(actionGroup, toggleActionEntries)

	// add radio actions
	addRadioActions(actionGroup, colorActionEntries, colorRed, nil)

	// add radio actions
	addRadioActions(actionGroup, shapeActionEntries, shapeSquare, nil)

	merge := gtk.NewUIManager()
	merge.InsertActionGroup(actionGroup, 0)
	window.AddAccelGroup(merge.GetAccelGroup())

	merge.AddUiFromString(uiInfo, int64(len(uiInfo)))
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
	infoBar.SetNoShowAll(true)
	messageLabel := gtk.NewLabel("")
	messageLabel.Show()

	//infoBar.GetContentArea().
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
		// update status bar
	})
	buffer.Connect(gtk.SigMarkSet, func() {
		// mark_set_callback
	})

	updateStatusBar(buffer, statusBar)
	window.SetDefaultSize(200, 200)
	window.ShowAll()
	gtk.Main()
}

func updateStatusBar(buffer gtk.TextBuffer, statusBar gtk.Statusbar) {
	// TODO
}
