package main

import (
	"log"
	"unsafe"

	"github.com/linuxdeepin/go-gir/g-2.0"

	"github.com/linuxdeepin/go-gir/gudev-1.0"
	"github.com/linuxdeepin/go-gir/gi"
)

func main() {
	arr := gi.NewCStrArrayZTWithStrings("input")
	client := gudev.NewClient(arr)
	arr.FreeAll()

	devices := client.QueryBySubsystem("input")
	devices.ForEach(func(item unsafe.Pointer) {
		dev := gudev.WrapDevice(item)
		name := dev.GetName()
		log.Println(name)
	})

	devices.FreeFull(func(item unsafe.Pointer) {
		obj := g.WrapObject(item)
		obj.Unref()
	})
}
