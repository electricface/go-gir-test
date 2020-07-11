package main

/*
void* getPointer_myFunc();
void callIt(void * fn);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export myFunc
func myFunc() int {
	fmt.Println("call myFunc")
	return 0
}

func GetPointer_myFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myFunc())
}

func main() {
	p := unsafe.Pointer(C.getPointer_myFunc())
	fmt.Println(p)
	C.callIt(p)
}
