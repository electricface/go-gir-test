package main

import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir/g-2.0"
)

func main() {

	list := g.NewList()
	list = list.Append(unsafe.Pointer(uintptr(1)))
	list = list.Append(unsafe.Pointer(uintptr(2)))
	list = list.Append(unsafe.Pointer(uintptr(3)))
	list = list.Append(unsafe.Pointer(uintptr(4)))

	list.ForEachC(func(args interface{}) {
		a := args.(*g.FuncStruct)
		log.Println(uintptr(a.F_data))
	})
}
