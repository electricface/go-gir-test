package main

import (
	"fmt"

	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	cal := gtk.NewCalendar()
	cal.SetDetailFunc(func(calendar gtk.Calendar, year uint32, month uint32, day uint32) (result string) {
		//log.Println(year, month, day)
		return fmt.Sprintf("%v %v %v", year, month, day)
	})
	cal.SetDetailFunc(nil)
	cal.SetDetailFunc(func(calendar gtk.Calendar, year uint32, month uint32, day uint32) (result string) {
		//log.Println(year, month, day)
		return fmt.Sprintf("%v %v %v", year, month, day)
	})
	win.Add(cal)
	win.ShowAll()
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	gtk.Main()
}
