package main

import (
	"log"
	"unsafe"

	"github.com/linuxdeepin/go-gir/gi"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	stockIds := gtk.StockListIds()
	stockIds.ForEach(func(item unsafe.Pointer) {
		stockId := gi.GoString(item)
		log.Println(stockId)
	})
	stockIds.FreeFull(func(item unsafe.Pointer) {
		gi.Free(item)
	})

	stockIds = gtk.StockListIds()
	stockIds.ForEachC(func(v interface{}) {
		args := v.(*g.FuncStruct)
		stockId := gi.GoString(args.F_data)
		log.Println(stockId)
	})
	stockIds.FreeFull(func(item unsafe.Pointer) {
		gi.Free(item)
	})

}
