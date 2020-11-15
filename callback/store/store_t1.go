package main

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"

	"github.com/linuxdeepin/go-gir/gdk-3.0"
	"github.com/linuxdeepin/go-gir/gi"

	"github.com/davecgh/go-spew/spew"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	w := gtk.NewWindow(gtk.WindowTypeToplevel)
	w.Show()
	w.SetSizeRequest(100, 100)
	w.Connect(gtk.SigDestroy, gtk.MainQuit)
	w.Connect(gtk.SigButtonPressEvent, func(args []interface{}) interface{} {
		log.Println("button press")
		spew.Dump(args)
		var arg0 gtk.Window
		var ev gdk.Event
		gi.Store(args, &arg0, &ev)
		log.Printf("arg0: %#v, arg1: %#v\n", arg0, ev)
		_, btn := ev.GetButton()
		_, x, y := ev.GetCoords()
		log.Println(x, y)
		log.Println(btn)

		var args0 struct {
			Window gtk.Window
			Ev     gdk.Event
		}
		gi.StoreStruct(args, &args0)
		//spew.Dump(args0)
		//args0.Window.SetSizeRequest(300, 300)
		ev = args0.Ev
		_, btn = ev.GetButton()
		_, x, y = ev.GetCoords()
		args0.Window.SetTitle(fmt.Sprint(btn, x, y))
		log.Println("- ", x, y)
		log.Println("- ", btn)
		return true
	})
	gtk.Main()
}

func StoreStruct(args []interface{}, dest interface{}) {
	rv := reflect.ValueOf(dest)
	if rv.Kind() == reflect.Ptr {
		log.Println("is ptr")
		elem := rv.Elem()
		if elem.Kind() == reflect.Struct {
			num := elem.NumField()
			for i := 0; i < num; i++ {
				if i >= len(args) {
					break
				}
				src := args[i]
				f := elem.Field(i)
				store(src, f.Addr().Interface())
			}
		}
	}
}

func Store(args []interface{}, destSlice ...interface{}) {
	for i, arg := range args {
		if i >= len(destSlice) {
			break
		}
		dest := destSlice[i]
		store(arg, dest)
	}
}

func store(src interface{}, dest interface{}) {
	switch a := src.(type) {
	//case g.Object:
	//	left, ok := dest.(*g.Object)
	//	if ok {
	//		left.P = a.P
	//	} else {
	//		storeStructFieldP(dest, a.P)
	//	}
	case unsafe.Pointer:
		left, ok := dest.(*unsafe.Pointer)
		if ok {
			*left = a
		} else {
			storeStructFieldP(dest, a)
		}
	default:
		srcRv := reflect.ValueOf(src)
		if srcRv.Kind() == reflect.Struct {
			p := srcRv.FieldByName("P")
			if p.Kind() == reflect.UnsafePointer {
				storeStructFieldP(dest, unsafe.Pointer(p.Pointer()))
			}
		}

	}
}

func storeStructFieldP(dest interface{}, ptr unsafe.Pointer) {
	rv := reflect.ValueOf(dest)
	if rv.Kind() == reflect.Ptr {
		elem := rv.Elem()
		if elem.Kind() == reflect.Struct {
			p := elem.FieldByName("P")
			if p.IsValid() && p.Kind() == reflect.UnsafePointer {
				p.SetPointer(ptr)
			}
		}
	}
}
