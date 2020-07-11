package main

import "C"
import (
	"unsafe"

	"github.com/linuxdeepin/go-gir/gi"
)

func main() {

	p := C.CString("hello world")
	gi.Free(unsafe.Pointer(p))
	// 在普通情况下，gi.Free 调用的 g_free 就是 free。
	gi.Free(unsafe.Pointer(p))
}
