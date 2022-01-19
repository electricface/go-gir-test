package main

import (
	"log"

	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	builder := gtk.NewBuilder()
	_, err := builder.AddFromFile("pygobject-demos/data/demo.ui")
	if err != nil {
		log.Fatal(err)
	}
	window := gtk.WrapWindow(builder.GetObject("window1").P)
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	actAbout := gtk.WrapAction(builder.GetObject("About").P)
	actAbout.Connect(gtk.SigActivate, func() {
		// about_activate
		aboutDlg := gtk.WrapAboutDialog(builder.GetObject("aboutdialog1").P)
		aboutDlg.Run()
		aboutDlg.Hide()
	})

	actQuit := gtk.WrapAction(builder.GetObject("Quit").P)
	actQuit.Connect(gtk.SigActivate, func() {
		gtk.MainQuit()
	})

	window.ShowAll()
	gtk.Main()
}
