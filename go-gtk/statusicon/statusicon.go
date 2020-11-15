package main

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
extern void my_menu_position_fn(GtkMenu *menu,
                       gint *x,
                       gint *y,
                       gboolean *push_in,
                       gpointer user_data);


static void* getPointer() {
	return (void*)(my_menu_position_fn);
}

*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

//export my_menu_position_fn
func my_menu_position_fn(menu *C.GtkMenu, x *C.gint, y *C.gint, pushIn *C.gboolean, userData C.gpointer) {
	fmt.Println("call export c")
	call := callMap[uint(uintptr(userData))]

	args := &Struct1{
		Menu:    gtk.WrapMenu(unsafe.Pointer(menu)),
		XP:      unsafe.Pointer(x),
		YP:      unsafe.Pointer(y),
		PushInP: unsafe.Pointer(pushIn),
	}
	call(args)
}

func getPointer() unsafe.Pointer {
	return C.getPointer()
}

type Struct1 struct {
	Menu    gtk.Menu
	XP      unsafe.Pointer
	YP      unsafe.Pointer
	PushInP unsafe.Pointer
}

func myCall2(args interface{}) {
	s1 := args.(*Struct1)

	log.Println(s1.Menu.P)

	x := (*C.gint)(s1.XP)
	valX := *x
	log.Println("valX:", valX)
	*x = 0

	y := (*C.gint)(s1.YP)
	valY := *y
	log.Println("valY:", valY)
	*y = 0

	pushIn := (*C.gboolean)(s1.PushInP)
	valPushIn := *pushIn
	log.Println("pushIn:", valPushIn)
}

func myCall3(args interface{}) {
	s1 := args.(*gtk.MenuPositionFuncStruct)

	log.Println("menu ptr is", s1.F_menu.P)
	x := (*int32)(s1.F_x)
	*x = 0

	y := (*int32)(s1.F_y)
	*y = 0
	//log.Println(s1.Menu.P)
	//
	//x := (*C.gint)(s1.XP)
	//valX := *x
	//log.Println("valX:", valX)
	//*x = 0
	//
	//y := (*C.gint)(s1.YP)
	//valY := *y
	//log.Println("valY:", valY)
	//*y = 0
	//
	//pushIn := (*C.gboolean)(s1.PushInP)
	//valPushIn := *pushIn
	//log.Println("pushIn:", valPushIn)
}

//func myCall1(args []interface{}) {
//	var menu gtk.Menu
//	var xP unsafe.Pointer
//	var yP unsafe.Pointer
//	var pushInP unsafe.Pointer
//	err := dbus.Store(args, &menu, &xP, &yP, &pushInP)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(menu.P)
//
//	x := (*C.gint)(xP)
//	valX := *x
//	log.Println("valX:", valX)
//
//	y := (*C.gint)(yP)
//	valY := *y
//	log.Println("valY:", valY)
//
//	pushIn := (*C.gboolean)(pushInP)
//	valPushIn := *pushIn
//	log.Println("pushIn:", valPushIn)
//}

func myCall(args []interface{}) {
	fmt.Println("call myCall")
	spew.Dump(args)
	menu := args[0].(gtk.Menu)
	xP := args[1].(unsafe.Pointer)
	yP := args[2].(unsafe.Pointer)
	pushInP := args[3].(unsafe.Pointer)

	log.Println(menu.P)

	x := (*C.gint)(xP)
	valX := *x
	log.Println("valX:", valX)

	y := (*C.gint)(yP)
	valY := *y
	log.Println("valY:", valY)

	pushIn := (*C.gboolean)(pushInP)
	valPushIn := *pushIn
	log.Println("pushIn:", valPushIn)
}

var callMap = make(map[uint]func(args interface{}))

func main() {
	//ptr := getPointer()
	//log.Println(ptr)

	gtk.Init(0, 0)

	g.SetApplicationName("go-gtk-statusicon-example")

	menuItemQuit := gtk.NewMenuItemWithLabel("Quit")
	menuItemQuit.Connect(gtk.SigActivate, func() {
		gtk.MainQuit()
	})
	menu := gtk.NewMenu()
	menu.Append(menuItemQuit)
	menu.ShowAll()
	log.Println("menu.P is", menu.P)

	si := gtk.NewStatusIconFromStock(gtk.STOCK_FILE)
	si.SetTitle("StatusIcon Example")
	si.SetTooltipMarkup("StatusIcon Example")
	si.Connect(gtk.SigPopupMenu, func(args []interface{}) {
		button := args[1].(uint32)
		activateTime := args[2].(uint32)
		// TODO 这里 func 传了 0，就是空指针，也是可以用的。

		menu.Popup(nil, nil, myCall3, button, activateTime)
		log.Println(button, activateTime)
	})

	//	fmt.Println(`
	//Can you see statusicon in systray?
	//If you don't see it and if you use 'unity', try following.
	//
	//# gsettings set com.canonical.Unity.Panel systray-whitelist \
	//  "$(gsettings get com.canonical.Unity.Panel systray-whitelist \|
	//  sed -e "s/]$/, 'go-gtk-statusicon-example']/")"
	//`)

	gtk.Main()
}
