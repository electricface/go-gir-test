package main

import "C"
import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gtk-3.0"
)

func menuPositionFn(menu gtk.Menu, xP unsafe.Pointer, yP unsafe.Pointer) (pushIn bool) {
	log.Println(menu.P)

	x := (*C.gint)(xP)
	valX := *x
	log.Println("valX:", valX)

	y := (*C.gint)(yP)
	valY := *y
	log.Println("valY:", valY)

	*x = 0
	*y = 0
	return false
}

func main() {
	gtk.Init(0, 0)

	g.SetApplicationName("go-gtk-statusicon-example")
	menuItemQuit := gtk.NewMenuItemWithLabel("Quit")
	menuItemQuit.Connect(gtk.SigActivate, func() {
		gtk.MainQuit()
	})
	menu := gtk.NewMenu()
	menu.Append(menuItemQuit)
	menu.ShowAll()
	log.Println("menu.P is", menu.P)

	si := gtk.NewStatusIconFromStock(gtk.STOCK_FILE)
	si.SetTitle("StatusIcon Example")
	si.SetTooltipMarkup("StatusIcon Example")
	si.Connect(gtk.SigPopupMenu, func(args []interface{}) {
		button := args[1].(uint32)
		activateTime := args[2].(uint32)
		menu.Popup(nil, nil, menuPositionFn, button, activateTime)
		log.Println(button, activateTime)
	})
	gtk.Main()
}
