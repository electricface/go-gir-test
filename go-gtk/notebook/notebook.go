package main

import (
	"strconv"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("GTK Notebook")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	notebook := gtk.NewNotebook()
	for n := 1; n <= 10; n++ {
		page := gtk.NewFrame("demo" + strconv.Itoa(n))
		notebook.AppendPage(page, gtk.NewLabel("demo"+strconv.Itoa(n)))

		hBox := gtk.NewHBox(false, 1)

		prev := gtk.NewButtonWithLabel("go prev")
		prev.Connect(gtk.SigClicked, func() {
			notebook.PrevPage()
		})
		hBox.Add(prev)

		next := gtk.NewButtonWithLabel("go next")
		next.Connect(gtk.SigClicked, func() {
			notebook.NextPage()
		})
		hBox.Add(next)

		page.Add(hBox)
	}

	window.Add(notebook)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
