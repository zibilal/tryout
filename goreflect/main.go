package main

import (
	"reflect"
	"fmt"
	"io"
	"os"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}
