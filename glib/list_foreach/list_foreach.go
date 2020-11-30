package main

import (
	"log"
	"unsafe"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
)

func main() {
	l := g.NewList()
	l.Append(gi.Uint2Ptr(1))
	l.Append(gi.Uint2Ptr(2))
	l.Append(gi.Uint2Ptr(3))
	l.Append(gi.Uint2Ptr(4))

	l.ForEach(func(item unsafe.Pointer) {
		log.Println(item)
	})
	l.ForEachC(func(data unsafe.Pointer) {
		log.Println("data:", data)
	})
}
