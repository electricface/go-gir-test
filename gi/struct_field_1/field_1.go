package main

/*
#cgo pkg-config: pango
#define PANGO_ENABLE_ENGINE
#define PANGO_ENABLE_BACKEND
#include <pango/pango-modules.h>
#include <pango/pango.h>
*/
import "C"
import (
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/linuxdeepin/go-gir/gi"
)

// Struct AttrSize
type AttrSize struct {
	P unsafe.Pointer
}

const SizeOfStructAttrSize = 24

//func AttrSizeGetType() gi.GType {
//	ret := _I.GetGType(12, "AttrSize")
//	return ret
//}
func (v AttrSize) p() *C.PangoAttrSize {
	return (*C.PangoAttrSize)(v.P)
}

// struct field attr
func (v AttrSize) Attr() (result int /*TODO*/) {
	return
}

// struct field size
func (v AttrSize) Size() (result int32) {
	result = int32(v.p().size)
	return
}

// struct field absolute
//func (v AttrSize) Absolute() (result uint32) {
//	result = uint32(v.p().absolute)
//	return
//}

func main() {
	v := AttrSize{P: gi.SliceAlloc0(SizeOfStructAttrSize)}
	spew.Dump(*v.p())
}
