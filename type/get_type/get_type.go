package main

import (
	"fmt"

	"github.com/linuxdeepin/go-gir/gtk-3.0"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gst-1.0"
)

func main() {
	// 以下三个 gst 的 GetType 都不是用 _I.GetGType 获取的，而是调用 xxx_get_type 方法获取的。
	v := gst.StaticCapsGetType()
	fmt.Println("static caps type:", v)

	v = gst.StaticPadTemplateGetType()
	fmt.Println("static pad template type:", v)

	v = gst.TypeFindGetType()
	fmt.Println("type find type:", v)

	v = gtk.WidgetGetType()
	fmt.Println("widget type:", v)

	v = g.FileInfoGetType()
	fmt.Println("file info type:", v)

	v = g.SettingsGetType()
	fmt.Println("settings type:", v)

	v = gst.BinGetType()
	fmt.Println("gst.Bin type:", v)

}
