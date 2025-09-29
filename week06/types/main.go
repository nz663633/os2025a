package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f64 float64
	var str string
	var i32 int32
	var b bool
	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(i32, reflect.TypeOf(i32))
	fmt.Println(b, reflect.TypeOf(b))
}