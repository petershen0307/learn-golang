package main

/*
#cgo LDFLAGS: -ldl
#cgo CFLAGS: -I ../lib-c-code
#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>
#include "main.h"
void wrapper(void* func_p, int a) {
	hello_world_t f = (hello_world_t)func_p;
	f(a);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	shared_lib_path := C.CString("./main.so")
	defer C.free(unsafe.Pointer(shared_lib_path))
	lib_main := C.dlopen(shared_lib_path, C.RTLD_LAZY)
	if e := C.dlerror(); e != nil {
		fmt.Println("error", C.GoString(e))
		return
	}
	defer C.dlclose(lib_main)
	symbolStr := C.CString("hello_world")
	defer C.free(unsafe.Pointer(symbolStr))
	func_p := C.dlsym(lib_main, symbolStr)
	if e := C.dlerror(); e != nil {
		fmt.Println("error", C.GoString(e))
		return
	}
	// https://forum.golangbridge.org/t/how-to-call-shared-object-function-loaded-dynamically-from-pointer-in-go/33252
	C.wrapper(func_p, C.int(42))
}
