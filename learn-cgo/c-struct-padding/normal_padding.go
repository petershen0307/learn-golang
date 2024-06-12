package main

/*
#include <stdio.h>
typedef struct {
	int a; // 4 bytes + 4 bytes padding
	void* p; // 8 bytes
} MyStruct;

static inline void print_MyStruct(){
	printf("c normalPadding %ld\n", sizeof(MyStruct));
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func normalPadding() {
	C.print_MyStruct()
	// the size of C.MyStruct is also 16 = 4 + 4 + 8, the structure size is the same of C sizeof
	a := C.MyStruct{}
	a.a = C.int(1)
	fmt.Println("go normalPadding", unsafe.Sizeof(a))
}
