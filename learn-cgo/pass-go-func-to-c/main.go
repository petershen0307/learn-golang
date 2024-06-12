package main

/*
typedef void (*callback) (int);
extern void hello_world(int a);
static inline void wrapper(callback f, int a) {
	f(a);
}
*/
import "C"
import "fmt"

//export hello_world
func hello_world(a C.int) {
	fmt.Println(a)
}
func main() {
	C.wrapper(C.callback(C.hello_world), C.int(42))
}
