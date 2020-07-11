package main

/*
#include <glib-object.h>
#cgo pkg-config: glib-2.0
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GType C.GType

func main() {
	gTypeSize := unsafe.Sizeof(GType(0))
	fmt.Println("", gTypeSize)
}
