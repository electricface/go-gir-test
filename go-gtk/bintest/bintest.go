package main

import (
	"fmt"
	"log"

	//dbus "pkg.deepin.io/lib/dbus1"

	"github.com/electricface/go-gir/gtk-3.0"

	"reflect"
	//"pkg.deepin.io/lib/dbusutil"
)

var _ reflect.Type

type structA struct {
}

func (a *structA) GetInterfaceName() string {
	return "user.swt.structA"
}

func (a *structA) Public() {
	fmt.Println("hello world")
}

func main() {
	func1 := func() {
		fmt.Println("hello world")
	}
	val := reflect.ValueOf(func1)
	ifc := val.Interface()
	log.Println(ifc)
	fn := ifc.(func())
	fn()
	//val.Call(nil)

	//bus, err := dbus.SessionBus()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//service := dbusutil.NewService(bus)
	//a := &structA{}
	//service.Export("/a", a)

	gtk.Init(0, 0)
	w := gtk.NewWindow(gtk.WindowTypeToplevel)
	w.SetTitle("hello world")
	gtk.Main()
}
