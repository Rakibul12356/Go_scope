package main

import (
	"fmt"
	"hello/mathlib"
)

func main() {
	a := 20
	fmt.Println(a)
	add()
	result := mathlib.Money
	fmt.Println(result)
	sum := mathlib.Sum(10, 20)
	fmt.Println(sum)
}
