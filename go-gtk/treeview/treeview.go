package main

import (
	//"strconv"

	"log"
	"strconv"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gdkpixbuf-2.0"
	gi "github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("GTK Folder View")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)

	typeArr := gi.NewGTypeArray(gdkpixbuf.PixbufGetType(), g.TYPE_STRING)
	store := gtk.NewTreeStore(2, typeArr)
	defer typeArr.Free()

	treeView := gtk.NewTreeView()
	swin.Add(treeView)

	treeView.SetModel(store)
	treeView.AppendColumn(gtk.NewTreeViewColumnWithAttribute("pixbuf", gtk.NewCellRendererPixbuf(),
		"pixbuf", 0))
	treeView.AppendColumn(gtk.NewTreeViewColumnWithAttribute("text", gtk.NewCellRendererText(), "text", 1))

	iter1 := gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}
	iter2 := gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}
	iter3 := gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}

	for n := 1; n <= 10; n++ {
		store.Append(iter1, gtk.TreeIter{})
		p1 := gtk.NewImage().RenderIcon(gtk.STOCK_DIRECTORY, int32(gtk.IconSizeSmallToolbar), "")
		v1 := g.NewValueT(gdkpixbuf.PixbufGetType())
		v1.SetObject(p1)
		store.SetValue(iter1, 0, v1)

		v2, err := g.NewValueWith("Folder" + strconv.Itoa(n))
		if err != nil {
			log.Fatal(err)
		}
		store.SetValue(iter1, 1, v2)

		store.Append(iter2, iter1)
		p1 = gtk.NewImage().RenderIcon(gtk.STOCK_DIRECTORY, int32(gtk.IconSizeSmallToolbar), "")
		v1 = g.NewValueT(gdkpixbuf.PixbufGetType())
		v1.SetObject(p1)
		store.SetValue(iter2, 0, v1)

		v2, err = g.NewValueWith("SubFolder" + strconv.Itoa(n))
		if err != nil {
			log.Fatal(err)
		}
		store.SetValue(iter2, 1, v2)

		store.Append(iter3, iter2)
		p1 = gtk.NewImage().RenderIcon(gtk.STOCK_FILE, int32(gtk.IconSizeSmallToolbar), "")
		v1 = g.NewValueT(gdkpixbuf.PixbufGetType())
		v1.SetObject(p1)
		store.SetValue(iter3, 0, v1)

		v2, err = g.NewValueWith("File" + strconv.Itoa(n))
		if err != nil {
			log.Fatal(err)
		}
		store.SetValue(iter3, 1, v2)

	}
	iter1.Free()
	iter2.Free()
	iter3.Free()

	//treeview.Connect("row_activated", func() {
	//	var path *gtk.TreePath
	//	var column *gtk.TreeViewColumn
	//	treeview.GetCursor(&path, &column)
	//	mes := "TreePath is: " + path.String()
	//	dialog := gtk.NewMessageDialog(
	//		treeview.GetTopLevelAsWindow(),
	//		gtk.DIALOG_MODAL,
	//		gtk.MESSAGE_INFO,
	//		gtk.BUTTONS_OK,
	//		mes)
	//	dialog.SetTitle("TreePath")
	//	dialog.Response(func() {
	//		dialog.Destroy()
	//	})
	//	dialog.Run()
	//})
	treeView.Connect(gtk.SigRowActivated, func() {
		log.Println("signal row_activated")
		path, column := treeView.GetCursor()
		mes := "TreePath is:" + path.ToString()
		log.Println(mes)
		_ = column
		//column.GetProperty()
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
