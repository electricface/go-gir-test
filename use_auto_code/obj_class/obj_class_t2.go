package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gtk-3.0"
)

func main3() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	objClass := win.GetClass()
	widgetClass := gtk.WidgetClass{P: objClass.P}

	cssName := widgetClass.GetCssName()
	log.Println("cssName:", cssName)
}

func init() {
	log.Println("run")
}

func main() {
	winGType := gtk.WindowGetType()
	typeClass := g.TypeClassRef(winGType)
	widgetClass := gtk.WidgetClass{P: typeClass.P}
	cssName := widgetClass.GetCssName()
	log.Println("cssName:", cssName)
}
