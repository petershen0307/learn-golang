package main

import "fmt"

func main() {
	var a func() int
	{
		i := 10
		a = foo(&i)
	}
	fmt.Println(a())
}

func foo(i *int) func() int {
	fmt.Printf("foo:%x\n", &i)
	return func() int {
		fmt.Printf("anonymous:%x\n", &i)
		return *i + 100
	}
}
