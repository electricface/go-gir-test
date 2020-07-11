package main

import (
	"unsafe"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gdkpixbuf-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
	gi "github.com/linuxdeepin/go-gir/gi"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("GTK Icon View")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)

	typeArr := gi.NewGTypeArray(gdkpixbuf.PixbufGetType(), g.TYPE_STRING)
	store := gtk.NewListStore(2, typeArr)
	typeArr.Free()
	iconview := gtk.NewIconViewWithModel(store)
	iconview.SetPixbufColumn(0)
	iconview.SetTextColumn(1)
	swin.Add(iconview)

	stockIdList := gtk.StockListIds()

	iterMem := gi.Malloc(gtk.SizeOfStructTreeIter)
	iter := gtk.TreeIter{P: iterMem}
	// TODO: iter := gtk.AllocTreeIter()

	stockIdList.ForEach(func(item unsafe.Pointer) {
		store.Append(iter)

		str := gi.GoString(item)
		pixBuf := gtk.NewImage().RenderIcon(str, int32(gtk.IconSizeSmallToolbar), "")

		v1, _ := g.NewValueWith(pixBuf.Object)
		store.SetValue(iter, 0, v1)

		v2, _ := g.NewValueWith(str)
		store.SetValue(iter, 1, v2)
	})
	iter.Free()

	// go-gtk 中的写法：
	//gtk.StockListIDs().ForEach(func(d unsafe.Pointer, v interface{}) {
	//	id := glib.GPtrToString(d)
	//	var iter gtk.TreeIter
	//	store.Append(&iter)
	//	store.Set(&iter,
	//		0, gtk.NewImage().RenderIcon(id, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf,
	//		1, id)
	//})

	window.Add(swin)
	window.SetSizeRequest(500, 200)
	window.ShowAll()

	gtk.Main()
}
