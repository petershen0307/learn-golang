package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	type foo struct {
		A string
		B string
		C string
	}

	f := foo{}
	ff := reflect.ValueOf(&f).Elem()
	for i := 0; i != ff.NumField(); i++ {
		v := ff.Field(i)
		fmt.Println(v.CanSet())
		v.SetString(fmt.Sprint(i))
		fmt.Println(v.Type(), ":", v.String())
	}
}

func Test_map(t *testing.T) {
	m := map[string]string{
		"a": "1",
		"b": "2",
	}
	iter := reflect.ValueOf(m).MapRange()
	// panic if call Key() before Next()
	// panic: MapIter.Key called before Next
	// k := iter.Key()
	// v := iter.Value()
	// fmt.Println(k, ":", v)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Println(k, ":", v)
	}
}
