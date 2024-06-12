package main

/*
#include <stdio.h>
typedef struct {
	char b; // 1 bytes
	char z[]; // sizeof will omit this field
} char_flexible_char;

typedef struct {
	char b; // 1 bytes + 3 bytes padding
	int z[]; // sizeof will omit this field
} char_flexible_int;

typedef struct {
	int b; // 4 bytes
	char z[]; // sizeof will omit this field
} int_flexible_char;

typedef struct {
	int b; // 4 bytes
	int z[]; // sizeof will omit this field
} int_flexible_int;

static inline void print_FlexibleArrayMember(){
	printf("c char_flexible_char %ld\n", sizeof(char_flexible_char));
	printf("c char_flexible_int %ld\n", sizeof(char_flexible_int));
	printf("c int_flexible_char %ld\n", sizeof(int_flexible_char));
	printf("c int_flexible_int %ld\n", sizeof(int_flexible_int));
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func flexibleArrayMember() {
	C.print_FlexibleArrayMember()
	// the size will be 1
	fmt.Println("go char_flexible_char", unsafe.Sizeof(C.char_flexible_char{}))
	// the size will be 1 + 3 padding!
	fmt.Println("go char_flexible_int", unsafe.Sizeof(C.char_flexible_int{}))
	// the size will be 4
	fmt.Println("go int_flexible_char", unsafe.Sizeof(C.int_flexible_char{}))
	// the size will be 4
	fmt.Println("go int_flexible_int", unsafe.Sizeof(C.int_flexible_int{}))
}
