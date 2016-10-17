package main

import (
	"reflect"
	"fmt"
)

func main(){
	square := func(x int) int{
		return x * x
	}
	fmt.Println("Starts...")
	data := []int {1, 2, 3, 4}
	result := Map(square, data)
	fmt.Println("The result", result)
}

func Map(f interface{}, xs interface{}) []interface{} {
	vf := reflect.ValueOf(f)
	vxs := reflect.ValueOf(xs)
	ys := make([]interface{}, vxs.Len())

	for i:=0; i < vxs.Len(); i++ {
		ys[i] = vf.Call([]reflect.Value{vxs.Index(i)})[0].Interface()
	}

	return ys
}
