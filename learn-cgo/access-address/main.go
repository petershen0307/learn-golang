package main

/*
#cgo CFLAGS: -Wpadded
#include <stdlib.h>
#include <stdio.h>
typedef struct {
    char number;
    int info[];
} MyStruct;
static inline MyStruct* get() {
	const char n = 2;
	void *p = malloc(sizeof(MyStruct) + n*sizeof(int));
	MyStruct *m = (MyStruct*)p;
	m->number = n;
	m->info[0] = 100;
	m->info[1] = 99;
	printf("c %p\n", m);
	// here MyStruct size is 4
	printf("c sizeof MyStruct %ld\n", sizeof(MyStruct));
	printf("c &m->number %p\n", &m->number);
	printf("c &m->info[0] %p\n", &m->info[0]);
	printf("c %d\n", m->number);
	printf("c %d\n", m->info[0]);
	printf("c %d\n", m->info[1]);
	return m;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := C.get()
	// here MyStruct size is 1, the padding is occurred
	fmt.Println("go sizeof MyStruct", unsafe.Sizeof(*a))
	fmt.Printf("go pointer %#x\n", uintptr(unsafe.Pointer(a)))
	b := unsafe.Slice((*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(a))+uintptr(4))), a.number)
	for _, v := range b {
		fmt.Println("go", v)
	}
	// want to access info
	for i := 0; i != int(a.number); i++ {
		b := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(4) + uintptr(i)*4))
		fmt.Println("go", b)
	}
}
