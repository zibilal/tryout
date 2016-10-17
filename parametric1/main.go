package main

import "fmt"

func main() {
	fmt.Println("Starts mapping...")

	square := func(x interface{}) interface{} {
		return x.(int) * x.(int)
	}

	nums := []int{1, 2, 3, 4}
	gnums := make([]interface{}, len(nums))
	for i, x := range nums {
		gnums[i] = x
	}

	result := Map(square, gnums)

	fmt.Println("Result", result)
}

func Map(f func(interface{}) interface{}, xs[]interface{}) [] interface{} {
	ys := make([]interface{}, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}
