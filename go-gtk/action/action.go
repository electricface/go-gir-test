package main

import (
	"fmt"
	"log"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := CreateWindow()
	window.SetPosition(gtk.WindowPositionCenter)
	window.ShowAll()
	gtk.Main()
}

func CreateWindow() gtk.Window {
	window := gtk.WrapWindow(gtk.NewWindow(gtk.WindowTypeToplevel).P)
	window.SetDefaultSize(700, 300)
	vbox := gtk.WrapVBox(gtk.NewVBox(false, 1).P)
	CreateMenuAndToolbar(window, vbox)
	CreateActivatableDemo(vbox)
	window.Add(vbox)
	window.Connect(gtk.SigDestroy, gtk.MainQuit)
	return window
}

func CreateActivatableDemo(vbox gtk.VBox) {
	action_entry := gtk.NewAction("ActionEntry",
		"Button attached to Action", "", gtk.STOCK_INFO)
	action_entry.Connect(gtk.SigActivate, func() {
		fmt.Println("Action clicked")
	})
	frame1 := gtk.WrapFrame(gtk.NewFrame("GtkActivatable interface demonstration").P)
	frame1.SetBorderWidth(5)
	hbox2 := gtk.WrapHBox(gtk.NewHBox(false, 5).P)
	hbox2.SetSizeRequest(400, 50)
	hbox2.SetBorderWidth(5)
	button1 := gtk.WrapButton(gtk.NewButton().P)
	button1.SetSizeRequest(250, 0)
	button1.SetRelatedAction(action_entry)
	hbox2.PackStart(button1, false, false, 0)
	hbox2.PackStart(gtk.NewVSeparator(), false, false, 0)
	button2 := gtk.NewButtonWithLabel("Hide Action")
	button2.SetSizeRequest(150, 0)
	button2.Connect(gtk.SigClicked, func() {
		action_entry.SetVisible(false)
		fmt.Println("Hide Action")
	})
	hbox2.PackStart(button2, false, false, 0)
	button3 := gtk.NewButtonWithLabel("Unhide Action")
	button3.SetSizeRequest(150, 0)
	button3.Connect(gtk.SigClicked, func() {
		action_entry.SetVisible(true)
		fmt.Println("Show Action")
	})
	hbox2.PackStart(button3, false, false, 0)
	frame1.Add(hbox2)
	vbox.PackStart(frame1, false, true, 0)
}

func CreateMenuAndToolbar(w gtk.Window, vbox gtk.VBox) {
	action_group := gtk.NewActionGroup("my_group")
	ui_manager := CreateUIManager()
	accel_group := ui_manager.GetAccelGroup()
	w.AddAccelGroup(accel_group)
	AddFileMenuActions(action_group)
	AddEditMenuActions(action_group)
	AddChoicesMenuActions(action_group)
	ui_manager.InsertActionGroup(action_group, 0)
	menubar := ui_manager.GetWidget("/MenuBar")
	vbox.PackStart(menubar, false, false, 0)
	toolbar := ui_manager.GetWidget("/ToolBar")
	vbox.PackStart(toolbar, false, false, 0)
	eventbox := gtk.NewEventBox()
	vbox.PackStart(eventbox, false, false, 0)
	label := gtk.NewLabel("Right-click to see the popup menu.")
	vbox.PackStart(label, false, false, 0)
}

func OnMenuFileNewGeneric() {
	fmt.Println("A File|New menu item was selected.")
}

func OnMenuFileQuit() {
	fmt.Println("quit app...")
	gtk.MainQuit()
}

func AddFileMenuActions(action_group gtk.ActionGroup) {
	action_group.AddAction(gtk.NewAction("FileMenu", "File", "", ""))

	action_filenewmenu := gtk.NewAction("FileNew", "New", "", gtk.STOCK_NEW)
	// TODO BUG 没有显示出 stock new 图标来
	action_group.AddAction(action_filenewmenu)

	action_new := gtk.NewAction("FileNewStandard", "_New",
		"Create a new file", gtk.STOCK_NEW)
	action_new.Connect(gtk.SigActivate, OnMenuFileNewGeneric)
	action_group.AddActionWithAccel(action_new, "")

	action_new_foo := gtk.NewAction("FileNewFoo", "New Foo",
		"Create new foo", gtk.STOCK_NEW)
	action_new_foo.Connect(gtk.SigActivate, OnMenuFileNewGeneric)
	action_group.AddAction(action_new_foo)

	action_new_goo := gtk.NewAction("FileNewGoo", "_New Goo",
		"Create new goo", gtk.STOCK_NEW)
	action_new_goo.Connect(gtk.SigActivate, OnMenuFileNewGeneric)
	action_group.AddAction(action_new_goo)

	action_filequit := gtk.NewAction("FileQuit", "quit hh", "", gtk.STOCK_QUIT)
	action_filequit.Connect(gtk.SigActivate, OnMenuFileQuit)
	action_group.AddActionWithAccel(action_filequit, "")
}

func OnMenuOther(args []interface{}) {
	obj := args[0].(g.Object)
	a := gtk.WrapAction(obj.P)
	actName := a.GetName()
	fmt.Println("on menu other", a.P)
	fmt.Println("action name:", actName)
}

func AddEditMenuActions(action_group gtk.ActionGroup) {
	action_group.AddAction(gtk.NewAction("EditMenu", "Edit", "", ""))

	action_editcopy := gtk.NewAction("EditCopy", "", "", gtk.STOCK_COPY)
	action_editcopy.Connect(gtk.SigActivate, OnMenuOther)
	action_group.AddActionWithAccel(action_editcopy, "")

	action_editpaste := gtk.NewAction("EditPaste", "", "", gtk.STOCK_PASTE)
	action_editpaste.Connect(gtk.SigActivate, OnMenuOther)
	action_group.AddActionWithAccel(action_editpaste, "")

	action_editsomething := gtk.NewAction("EditSomething", "Something", "", "")
	action_editsomething.Connect(gtk.SigActivate, OnMenuOther)
	action_group.AddActionWithAccel(action_editsomething, "<control><alt>S")
}

func AddChoicesMenuActions(action_group gtk.ActionGroup) {
	action_group.AddAction(gtk.NewAction("ChoicesMenu", "Choices", "", ""))

	var raSlice []gtk.RadioAction
	raOne := gtk.NewRadioAction("ChoiceOne", "One", "", "", 1)
	raOne.Connect(gtk.SigActivate, func() {
		cv := raOne.GetCurrentValue()
		log.Println("raOne activate", cv)
	})
	raSlice = append(raSlice, raOne)

	raTwo := gtk.NewRadioAction("ChoiceTwo", "Two", "", "", 2)
	raTwo.Connect(gtk.SigActivate, func() {
		cv := raTwo.GetCurrentValue()
		log.Println("raTwo activate", cv)
	})
	raSlice = append(raSlice, raTwo)

	raThree := gtk.NewRadioAction("ChoiceThree", "Three", "", "", 2)
	raThree.Connect(gtk.SigActivate, func() {
		cv := raThree.GetCurrentValue()
		log.Println("raThree activate", cv)
	})
	raSlice = append(raSlice, raThree)

	var sl g.SList
	for _, ra := range raSlice {
		ra.SetGroup(sl)
		sl = ra.GetGroup()
		action_group.AddAction(ra.Action)
	}

	ra_last := gtk.NewToggleAction("ChoiceToggle", "Toggle", "", "")
	ra_last.SetActive(true)
	action_group.AddAction(ra_last.Action)
}

func CreateUIManager() gtk.UIManager {
	UI_INFO := `
<ui>
  <menubar name='MenuBar'>
    <menu action='FileMenu'>
      <menu action='FileNew'>
        <menuitem action='FileNewStandard' />
        <menuitem action='FileNewFoo' />
        <menuitem action='FileNewGoo' />
      </menu>
      <separator />
      <menuitem action='FileQuit' />
    </menu>
    <menu action='EditMenu'>
      <menuitem action='EditCopy' />
      <menuitem action='EditPaste' />
      <menuitem action='EditSomething' />
    </menu>
    <menu action='ChoicesMenu'>
      <menuitem action='ChoiceOne'/>
      <menuitem action='ChoiceTwo'/>
      <menuitem action='ChoiceThree'/>
      <separator />
      <menuitem action='ChoiceToggle'/>
    </menu>
  </menubar>
  <toolbar name='ToolBar'>
    <toolitem action='FileNewStandard' />
    <toolitem action='FileQuit' />
  </toolbar>
  <popup name='PopupMenu'>
    <menuitem action='EditCopy' />
    <menuitem action='EditPaste' />
    <menuitem action='EditSomething' />
  </popup>
</ui>
`
	ui_manager := gtk.NewUIManager()
	_, err := ui_manager.AddUiFromString(UI_INFO, int64(len(UI_INFO)))
	if err != nil {
		log.Println(err)
	}
	return ui_manager
}
