package main 

import "fmt"

const a = 10 
 var p = 100 

func call (){
add := func (x int, y int) int {
	z := x + y
	fmt.Println("The sum is:", z)
	return z
}

add(5, 10)
add(a,p)
}

 func main(){
a:=20
	call()
	fmt.Println(a)

 }
 func init (){
	fmt.Println("init function called")
 }
