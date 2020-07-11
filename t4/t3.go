package main

/*
//static void callIt(void * fn) {
//    void (*a)();
//    a = fn;
//    a();
//}


extern void myFunc();
static void* getPointer_myFunc() {
	return (void*)(myFunc);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export myFunc
func myFunc() {
	fmt.Println("call myFunc")
}

func GetPointer_myFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myFunc())
}

func main() {
	//p := unsafe.Pointer(C.getPointer_myFunc())
	p := GetPointer_myFunc()
	fmt.Println(p)
	//C.callIt(p)
}
