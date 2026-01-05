package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //liststart
	y := x[:2] // 先頭から2つ。つまりx[0]とx[1]。残りもyのキャパシティに含まれている
	y = append(y, "z")
	y = append(y, "1")	
	y = append(y, "2")	
	fmt.Println("x:", x) // x: [a b z 1]
	fmt.Println("y:", y) // y: [a b z 1 2]  //listend
	fmt.Println(cap(x), cap(y)) // 4 8 
}
