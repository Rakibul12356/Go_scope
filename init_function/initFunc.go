package main

import "fmt"

var a int = 10

func main() {
	fmt.Println("I am main function")
	fmt.Println(a)
}

func init() {
	fmt.Println("I am init function", a)
	a = 20

}
