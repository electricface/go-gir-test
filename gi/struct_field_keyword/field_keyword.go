package main

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/electricface/go-gir/gi"
)

type Hook struct {
	P unsafe.Pointer
}

const SizeOfStructHook = 64

func (v Hook) p() *C.GHook {
	return (*C.GHook)(v.P)
}
func (v Hook) Data() (result unsafe.Pointer) {
	result = unsafe.Pointer(v.p().data)
	return
}
func (v Hook) Next() (result int /*TODO*/) {
	return
}
func (v Hook) Prev() (result int /*TODO*/) {
	return
}
func (v Hook) RefCount() (result uint32) {
	result = uint32(v.p().ref_count)
	return
}
func (v Hook) HookId() (result uint64) {
	result = uint64(v.p().hook_id)
	return
}
func (v Hook) Flags() (result uint32) {
	result = uint32(v.p().flags)
	return
}
func (v Hook) Func() (result unsafe.Pointer) {
	result = unsafe.Pointer(v.p()._func)
	return
}

func main() {
	h := Hook{P: gi.SliceAlloc0(SizeOfStructHook)}
	spew.Dump(*h.p())
}
