package main

import (
	"log"
	"testing"
	"unsafe"

	"pkg.deepin.io/gir/gio-2.0"

	"github.com/electricface/go-gir3/gi"
	"github.com/ying32/dylib"
)

var invokerCache = gi.NewInvokerCache("Gio")

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Gio", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}

}

const (
	File_new_for_path uint = iota
	File_get_uri
)

func FileNewForPath(path1 string) unsafe.Pointer {
	invoker, err := invokerCache.Get(File_new_for_path, "File", "new_for_path")
	if err != nil {
		log.Fatal(err)
	}

	var ret gi.Argument
	pPath := gi.CString(path1)
	arg1 := gi.NewStringArgument(pPath)
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	gi.Free(pPath)
	return ret.Pointer()
}

func FileGetPath(p unsafe.Pointer) string {
	invoker, err := invokerCache.Get(File_get_uri, "File", "get_uri")
	if err != nil {
		log.Fatal(err)
	}

	arg1 := gi.NewPointerArgument(p)
	var ret gi.Argument
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	return ret.String().Take()
}

func BenchmarkInvoke(b *testing.B) {
	fp := FileNewForPath("/home/tp1/hello-world")
	log.Printf("fp: %p\n", fp)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FileGetPath(fp)
	}
}

func BenchmarkGioCgo(b *testing.B) {
	//fp := FileNewForPath("/home/tp1/hello-world")
	//log.Printf("fp: %p\n", fp)
	//b.ResetTimer()
	//
	//for i := 0; i < b.N; i++ {
	//	FileGetPath(fp)
	//}
	f := gio.FileNewForPath("/home/tp1/hello-world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f.GetPath()
	}
}

func BenchmarkDylibCall(b *testing.B) {
	dll := dylib.NewLazyDLL("libgio-2.0.so")

	newForPath := dll.NewProc("g_file_new_for_path")

	path1 := "/home/tp1/hello-world"
	cPath1 := gi.CString(path1)
	fp, _, _ := newForPath.CallOriginal(uintptr(unsafe.Pointer(cPath1)))
	getPath := dll.NewProc("g_file_get_path")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		path, _, _ := getPath.CallOriginal(fp)
		arg := *(*gi.Argument)(unsafe.Pointer(&path))
		arg.String().Take()
	}
}

/*

BenchmarkInvoke-4        1001557          1318 ns/op          64 B/op          4 allocs/op with inline
BenchmarkDylibCall-4    1927447           684 ns/op          32 B/op          1 allocs/op

*/
