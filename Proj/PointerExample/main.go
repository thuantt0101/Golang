package main

import "fmt"

func main() {

	//1. ví dụ 1
	//Declare a Pointer
	//A pointer of type int
	var p1 *int
	fmt.Println("type of pointer p :", p1)
	//-->the result is type of pointer p : <nil>

	//2. ví dụ 2
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)            //7.98
	fmt.Println("address of a: ", &a) //x

	fmt.Println("p =  ", p)         //x
	fmt.Println("address of p", &p) //y

	fmt.Println("pp = ", pp) //y
	//Dereferencing a pointer to pointer

	fmt.Println("*pp", *pp)
	fmt.Println("**pp", **pp)
}
