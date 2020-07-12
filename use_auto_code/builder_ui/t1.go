package main

import (
	"fmt"
	"log"
	"os"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	app := gtk.NewApplication("org.rain.example", g.ApplicationFlagsFlagsNone)
	app.Connect(gtk.SigActivate, func(args []interface{}) {
		activate(args)
	})
	args := gi.NewCStrArrayWithStrings(os.Args...)
	defer args.FreeAll()
	status := app.Run(int32(len(os.Args)), args)
	app.Unref()
	os.Exit(int(status))
}

func activate(args []interface{}) {
	app := gtk.WrapApplication(args[0].(g.Object).P)
	builder := gtk.NewBuilder()
	_, err := builder.AddFromFile("./t1.ui")
	if err != nil {
		log.Fatal(err)
	}
	win := gtk.WrapWindow(builder.GetObject("window").P)
	app.AddWindow(win)

	button := gtk.WrapButton(builder.GetObject("print_hello").P)
	button.SetLabel("Print Hello")
	button.Connect(gtk.SigClicked, func() {
		fmt.Println("hello world")
	})

	entry := gtk.WrapEntry(builder.GetObject("entry").P)
	entry.SetEditable(true)
	entry.SetText("entry 001")
	button = gtk.WrapButton(builder.GetObject("print_entry").P)
	button.SetLabel("Print Entry")
	button.Connect(gtk.SigClicked, func() {
		entryTxt := entry.GetText()
		fmt.Println("entry:", entryTxt)
	})
}
