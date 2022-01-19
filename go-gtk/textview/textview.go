package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/electricface/go-gir3/gi"

	"github.com/electricface/go-gir/gdk-3.0"
	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)

	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetPosition(gtk.WindowPositionCenter)
	window.SetTitle("GTK Go!")
	window.SetIconName("textview")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	textview := gtk.NewTextView()
	textview.SetEditable(true)
	textview.SetCursorVisible(true)
	var iter gtk.TextIter
	iter.P = gi.Malloc(gtk.SizeOfStructTextIter)
	buffer := textview.GetBuffer()

	buffer.GetStartIter(iter)
	buffer.Insert(iter, "Hello ", int32(len("Hello ")))

	//buffer.CreateMark()
	//tag := buffer.CreateTag("bold", map[string]string{"background": "#FF0000", "weight": "700"})
	//buffer.InsertWithTag(&iter, "Google!", tag)

	//u := "http://www.google.com"
	//tag.SetData("tag-name", unsafe.Pointer(&u))
	//textview.Connect("event-after", func(ctx *glib.CallbackContext) {
	//	arg := ctx.Args(0)
	//	if ev := *(**gdk.EventAny)(unsafe.Pointer(&arg)); ev.Type != gdk.BUTTON_RELEASE {
	//		return
	//	}
	//	ev := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//	var iter gtk.TextIter
	//	textview.GetIterAtLocation(&iter, int(ev.X), int(ev.Y))
	//	tags := iter.GetTags()
	//	for n := uint(0); n < tags.Length(); n++ {
	//		vv := tags.NthData(n)
	//		tag := gtk.NewTextTagFromPointer(vv)
	//		u := *(*string)(tag.GetData("tag-name"))
	//		fmt.Println(u)
	//	}
	//})

	textview.Connect(gtk.SigEventAfter, func(args []interface{}) {
		evP := args[1].(unsafe.Pointer)
		ev := gdk.Event{P: evP}
		type0 := ev.GetEventType()
		if type0 != gdk.EventTypeButtonRelease {
			return
		}
		//evB := gdk.EventButton{P: evP}
		ok, x, y := ev.GetCoords()
		if ok {
			textview.GetIterAtLocation(iter, int32(x), int32(y))
			tags := iter.GetTags()
			tags.ForEach(func(item unsafe.Pointer) {
				tag := gtk.WrapTextTag(item)
				data := tag.GetData("tag-name")
				fmt.Println(data)
			})
		}

		log.Println(ev)
	})

	window.Add(textview)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
