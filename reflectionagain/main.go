package main

import (
	"fmt"
	"reflect"
	"io"
	"os"
	"time"
	"github.com/zibilal/tryout/reflectionagain/dstring"
)

func main() {
	fmt.Println("Starts...")

	t := reflect.TypeOf(3)
	fmt.Println("Type", t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))


	v := reflect.ValueOf(3) // a reflect.Value
	x1 := v.Interface() // an interface
	i := x1.(int) // an integer
	fmt.Printf("%d\n", i) // 3

	check(3)
	check(3.33)
	check(`Bilal Muhammad`)
	check(struct{first, last string}{"Bilal", "Muhammad"})

	var x int64 = 1

	var d time.Duration = 1 * time.Nanosecond

	fmt.Println(dstring.Any(x))
	fmt.Println(dstring.Any(d))
	fmt.Println(dstring.Any([]int64{x}))
	fmt.Println(dstring.Any([]time.Duration{d}))
	fmt.Println(dstring.Any(check))
}

func check(input interface{}) {
	v := reflect.ValueOf(input)
	fmt.Printf("%v is the value\n", v)
}
