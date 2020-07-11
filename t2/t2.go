package main

/*
#include <stdio.h>
void plusOne(int **i) {
    (**i)++;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//s1 := make([]*int, 5)
	var s1 [5]*int
	var a = 5
	s1[4] = &a
	pFirst := &s1[3]
	//C.plusOne((**C.int)((unsafe.Pointer)(pFirst)))
	plusOne4(pFirst)
	fmt.Println(s1[0])
}

func plusOne4(pFirst **int) {
	plusOne3(pFirst)
}

func plusOne3(pFirst **int) {
	plusOne2(pFirst)
}

func plusOne2(pFirst **int) {
	plusOne(pFirst)
}

func plusOne(pFirst **int) {
	C.plusOne((**C.int)((unsafe.Pointer)(pFirst)))
}
