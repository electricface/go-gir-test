package main

import "github.com/linuxdeepin/go-gir/gtk-3.0"

func main() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	win.Add(windowWidget())
	win.ShowAll()
	gtk.Main()
}

func windowWidget() gtk.Widget {
	grid := gtk.NewGrid()
	grid.SetOrientation(gtk.OrientationVertical)

	entry := gtk.NewEntry()
	label := gtk.NewLabel("hello world label")
	grid.Add(entry)
	entry.SetHexpand(true)
	grid.AttachNextTo(label, entry, gtk.PositionTypeRight, 1, 1)
	label.SetHexpand(true)

	entry.Connect(gtk.SigActivate, func() {
		s := entry.GetText()
		label.SetText(s)
	})

	sb := gtk.NewSpinButtonWithRange(0, 1, 0.1)
	pb := gtk.NewProgressBar()
	grid.Add(sb)
	sb.SetHexpand(true)
	grid.AttachNextTo(pb, sb, gtk.PositionTypeRight, 1, 1)
	pb.SetHexpand(true)
	sb.Connect(gtk.SigValueChanged, func() {
		pb.SetFraction(sb.GetValue())
	})

	label1 := gtk.NewLabel("")
	s := "Hyperlink to <a href=\"https://www.baidu.com/\">Baidu</a> for your clicking pleasure"
	label1.SetMarkup(s)
	grid.AttachNextTo(label1, sb, gtk.PositionTypeBottom, 2, 1)

	return grid.Widget
}
