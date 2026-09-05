package main

import "fmt"


// func processOperation (a int, b int,op func(p int , q int )){
// 	op(	a,b)
// }

// func add (x int, y int){
// 	z:=x+y
// 	fmt.Println("the result is ",z)

// }

// func main (){
// 	processOperation(10,20,add)

// }



func cell() func(x int, y int){
	
return add
}

func add(x int,y int){
	z:=x+y
	fmt.Println("the result is ",z)
}
func main (){
	sum:=cell()
	sum(10,20)
}