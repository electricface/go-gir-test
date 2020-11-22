package main

import (
	"fmt"
	"log"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	cal := gtk.NewCalendar()
	cal.SetDetailFunc(func(args *gtk.CalendarDetailFuncArgs) string {
		log.Println(args.Year, args.Month, args.Day)
		return fmt.Sprintf("%v %v %v", args.Year, args.Month, args.Day)
	})
	win.Add(cal)
	win.ShowAll()
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	gtk.Main()
}
