package main

import (
	"errors"
	"log"
	"strings"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gdkpixbuf-2.0"
	"github.com/electricface/go-gir/gi"
	"github.com/electricface/go-gir/gtk-3.0"
)

func newComboboxApp() {
	//demoApp :=
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("Combo boxes")
	window.SetBorderWidth(10)
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	vbox := gtk.NewVBox(false, 2)
	window.Add(vbox)

	frame := gtk.NewFrame("Some stock icons")
	vbox.PackStart(frame, false, false, 0)

	box := gtk.NewVBox(false, 0)
	box.SetBorderWidth(5)
	frame.Add(box)

	model := createStockIconStore()
	combo := gtk.NewComboBoxWithModel(model)
	box.Add(combo)
	rendererPb := gtk.NewCellRendererPixbuf()
	combo.PackStart(rendererPb, false)

	combo.AddAttribute(rendererPb, "pixbuf", 0)
	combo.SetCellDataFunc(rendererPb, setSensitive)

	rendererTxt := gtk.NewCellRendererText()
	combo.PackStart(rendererTxt, true)
	combo.AddAttribute(rendererTxt, "text", 1)
	combo.SetCellDataFunc(rendererTxt, setSensitive)
	combo.SetRowSeparatorFunc(func(model gtk.TreeModel, iter gtk.TreeIter) (result bool) {
		path := model.GetPath(iter)
		i := path.GetIndices()
		return i.AsSlice()[0] == 4
	})
	combo.SetActive(0)
	window.ShowAll()
}

func setSensitive(cell_layout gtk.CellLayout, cell gtk.CellRenderer, model gtk.TreeModel, iter gtk.TreeIter) {
	path := model.GetPath(iter)
	i := path.GetIndices()
	sensitive := !(i.AsSlice()[0] == 1)
	v, err := g.NewValueWith(sensitive)
	if err != nil {
		log.Fatal(err)
	}
	cell.SetProperty("sensitive", v)
	v.Free()
}

func createStockIconStore() gtk.ListStore {
	cellView := gtk.NewCellView()
	arr := gi.NewGTypeArray([]gi.GType{gdkpixbuf.PixbufGetType(), g.TYPE_STRING})
	store := gtk.NewListStore(2, arr)
	arr.Free()

	stockIds := []string{
		gtk.STOCK_DIALOG_WARNING,
		gtk.STOCK_STOP,
		gtk.STOCK_NEW,
		gtk.STOCK_CLEAR,
		"",
		gtk.STOCK_OPEN,
	}
	treeIter := gtk.TreeIter{P: gi.SliceAlloc0(gtk.SizeOfStructTreeIter)}
	for _, id := range stockIds {
		if id != "" {
			pixBuf := cellView.RenderIcon(id, int32(gtk.IconSizeButton), gi.NilStr)
			stockItem := gtk.StockItem{P: gi.SliceAlloc0(gtk.SizeOfStructStockItem)}
			ok := gtk.StockLookup(id, stockItem)
			label := ""
			if ok {
				label = strings.Replace(stockItem.Label(), "_", "", 1)
				log.Println("id:", id, "label:", label)
			} else {
				log.Fatal(errors.New("stock lookup failed"))
			}
			gi.SliceFree(gtk.SizeOfStructStockItem, stockItem.P)
			store.Append(treeIter)

			v0, err := g.NewValueT(gdkpixbuf.PixbufGetType())
			if err != nil {
				log.Fatal(err)
			}
			v0.SetObject(pixBuf)
			store.SetValue(treeIter, 0, v0)

			v1, err := g.NewValueWith(label)
			if err != nil {
				log.Fatal(err)
			}
			store.SetValue(treeIter, 1, v1)

			v0.Free()
			v1.Free()
		} else {
			store.Append(treeIter)
			store.SetValue(treeIter, 0, g.Value{})

			v1, err := g.NewValueWith("separator")
			if err != nil {
				log.Fatal(err)
			}
			store.SetValue(treeIter, 1, v1)
			store.SetValue(treeIter, 1, v1)
			v1.Free()
		}
	}
	gi.SliceFree(gtk.SizeOfStructTreeIter, treeIter.P)
	return store
}

func main() {
	log.SetFlags(log.Lshortfile)
	gtk.Init(0, 0)
	newComboboxApp()
	gtk.Main()
}
