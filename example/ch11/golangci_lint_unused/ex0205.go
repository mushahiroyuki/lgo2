package main

import "fmt"

func main() { //liststart
	x := 10 // このxの値は読まれることはない
	x = 20
	fmt.Println(x)
	x = 30 // このxの値も読まれることはない
} //listend

