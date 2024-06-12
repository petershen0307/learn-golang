package main

// https://stackoverflow.com/questions/35360023/how-to-link-a-library-by-its-exact-name-with-gcc
// LDFLAGS -lfoo will use the convention like libfoo.so
// LDFLAGS -l:foo.so but we still can specific shared library name

// use compile time linking, we need to set LD_LIBRARY_PATH environment when running the program
// export LD_LIBRARY_PATH=./learn-go
// we expect main.so under learn-cgo

/*
#cgo LDFLAGS: -L.. -l:main.so
#cgo CFLAGS: -I ../lib-c-code
#include <stdlib.h>
#include <dlfcn.h>
#include "main.h"
*/
import "C"

func main() {
	C.hello_world(42)
}
