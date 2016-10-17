package main

import (
	"reflect"
	"fmt"
)

func main() {
	fmt.Println("Starts mapping ...")
	result := Map(func(x int) int {return x * x}, []int{1, 2, 3, 4, 5})
	fmt.Println("Result", result)
}

func Map(f, xs interface{}) interface{} {
	vf := reflect.ValueOf(f)
	vxs := reflect.ValueOf(xs)

	tys := reflect.SliceOf(vf.Type().Out(0))
	vys := reflect.MakeSlice(tys, vxs.Len(), vxs.Len())

	for i:=0;i < vxs.Len(); i++ {
		y := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		vys.Index(i).Set(y)
	}
	return vys.Interface()
}


