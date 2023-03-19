package main

// #cgo LDFLAGS: -L. -lmain
// #include "c/main.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(FirstChar([]string{"hello", "noob", "e"}))
}

func FirstChar(list []string) string {
	cStrings := make([]*C.char, len(list))
	for i, s := range list {
		cStrings[i] = C.CString(s)
	}
	// Create a C array of pointers to the C strings
	cStringArray := (**C.char)(unsafe.Pointer(&cStrings[0]))
	res := C.FirstChar(cStringArray)
	str := C.GoString(res)
	for _, s := range cStrings {
		C.free(unsafe.Pointer(s))
	}
	return str
}