package main

import (
	"log"
	"testing"
	"unsafe"

	"github.com/linuxdeepin/go-gir/gi"
)

type DestroyNotify func(ptr unsafe.Pointer) int

func CallDestroyNotify(fn DestroyNotify, result unsafe.Pointer, args []unsafe.Pointer) {
	ptr := *(*unsafe.Pointer)(args[0])
	ret := fn(ptr)
	_ = ret
	//*(*int)(result) = ret
}

func BenchmarkT1(b *testing.B) {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GLib", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}

	aDN := func(ptr unsafe.Pointer) int {
		//log.Println("aDN", ptr)
		return 123
	}
	callableInfo := gi.GetCallableInfo("GLib", "DestroyNotify")
	defer callableInfo.Unref()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myTest(repo, callableInfo, aDN)
	}
}

func myTest(repo gi.Repository, callableInfo gi.CallableInfo, callback DestroyNotify) {

	cId, execPtr := gi.RegisterFClosure(func(result unsafe.Pointer, args []unsafe.Pointer) {
		CallDestroyNotify(callback, result, args)
	}, gi.ScopeCall, callableInfo)

	gi.CallMyDestroyFn(execPtr)
	gi.UnregisterFClosure(cId)
}
