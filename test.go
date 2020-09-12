package main

import (
	"fmt"
	"reflect"
)

func main() {
	//fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	var a int
	fmt.Println(reflect.TypeOf(a))
}
