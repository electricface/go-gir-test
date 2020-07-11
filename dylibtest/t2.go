package main

/*
#include <stdlib.h>
#include <stdio.h>
static void print_string(void* p) {
	printf("PS: %s", (char*)(p));
}
 */
import "C"
import (
	"github.com/ying32/dylib"
	//"github.com/ying32/govcl/vcl/api"
	"log"
	"reflect"
	"unsafe"
)

func strPtr(val *string) unsafe.Pointer {
	return unsafe.Pointer( (*reflect.StringHeader)(unsafe.Pointer(val)).Data )
}

type strPtrS struct {
	P unsafe.Pointer
}

func(v strPtrS) Take() string {
	str := C.GoString((*C.char)(v.P))
	//C.free(v.P)
	return str
	//arr := (*(*[1<<20]byte)(v.P))[:999999:999999]
	//for i, b := range arr {
	//	if b == 0 {
	//		return string(arr[:i])
	//		break
	//	}
	//}
	//C.print_string(v.P)
	//return copyStr(uintptr(v.P), 100)
	//return ""
}

// 这种跟copyStr3基本一样，只是用go来处理了
func copyStr(src uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	str := make([]uint8, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}



func main() {
	log.SetFlags(log.Lshortfile)

	glib := dylib.NewLazyDLL("libglib-2.0.so")
	t0, _, _ := glib.NewProc("g_date_time_new_now_local").Call()
	log.Println(t0)
	var format = "%x %X"
	strp := strPtr(&format)
	_ = strp
	strp2 := C.CString(format)
	result, _, _ := glib.NewProc("g_date_time_format").Call(t0, uintptr(unsafe.Pointer(strp2)))
	resS := strPtrS{ unsafe.Pointer(result)}.Take()
	log.Println(resS)

	dll := dylib.NewLazyDLL("libgio-2.0.so")

		newForPath := dll.NewProc("g_file_new_for_path")

		path1 := "http://tmp.com/startdde-login-sound-mark"
		cPath1 := C.CString(path1)
		fp, _, _ := newForPath.CallOriginal(uintptr(unsafe.Pointer(cPath1)))
		log.Printf("%p\n", unsafe.Pointer(fp))


		getUri := dll.NewProc("g_file_get_uri")
		uri, _, _ := getUri.CallOriginal(fp)
		log.Printf("%p\n", unsafe.Pointer(uri))
		uriStr := strPtrS{unsafe.Pointer(uri)}.Take()
		//uriStr := copyStr(uri, 100)
		log.Println(uriStr)

		getBasename := dll.NewProc("g_file_get_basename")
		bn, _, _ := getBasename.CallOriginal(fp)
		log.Printf("%p\n", unsafe.Pointer(bn))
		bnStr := strPtrS{unsafe.Pointer(bn)}.Take()
		log.Println(bnStr)

		isNative := dll.NewProc("g_file_is_native")
		isn, _, _ := isNative.CallOriginal(fp)
		log.Printf("%p\n", unsafe.Pointer(isn))
		log.Println(isn)

	//	vcl := dylib.NewLazyDLL("liblcl.so")
	//	err := vcl.Load()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//	appInst,_, _ := vcl.NewProc("Application_Instance").Call()
	//
	//	_,_,_ = vcl.NewProc("Application_Initialize").Call(appInst)
	//
	//	exeName, _, _ := vcl.NewProc("Application_GetExeName").Call(appInst)
	//
	//	exeNameStr := api.DStrToGoStr(exeName)
	//	log.Println(exeNameStr)
}
