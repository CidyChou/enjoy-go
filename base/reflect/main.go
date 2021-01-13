package main

import (
	"fmt"
	"reflect"
)

//T is T
type T struct{}

func main() {
	argRVs := []reflect.Value{}
	v := reflect.ValueOf(3)
	argRVs = append(argRVs, v)

	fmt.Println(argRVs[0])

	fmt.Println(v)

	fmt.Println(v == argRVs[0])

	s := v.Interface()
	fmt.Println(s)
	i := s.(int)
	fmt.Println(i)

	// fmt.Println()

	call()
}

func call() {
	name := "Do"
	name2 := "Do2"
	t := &T{}
	a := reflect.ValueOf(1111)
	b := reflect.ValueOf("world")
	in := []reflect.Value{a, b}
	reflect.ValueOf(t).MethodByName(name).Call(in)
	reflect.ValueOf(t).MethodByName(name2).Call(nil)
}

// Do is Do
func (t *T) Do(a int, b string) {
	fmt.Println("hello"+b, a)
}

// Do2 is Do
func (t *T) Do2() {
	fmt.Println("hello")
}
