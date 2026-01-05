package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //liststart
	y := x[:2] // 先頭から2つ。つまりx[0]とx[1]
	fmt.Println("1:", cap(x), cap(y)) // 1: 4 4
	y = append(y, "z")
	fmt.Println("2:", cap(x), cap(y)) // 2: 4 4
	y = append(y, "1")	
	fmt.Println("3:", cap(x), cap(y)) // 3: 4 4
	y = append(y, "2")	
	fmt.Println("4:", cap(x), cap(y)) // 4: 4 8
	fmt.Println("x:", x) // x: [a b z 1]
	fmt.Println("y:", y) // y: [a b z 1 2]
	fmt.Println(cap(x), cap(y)) // 4 8  //listend
}

/*
    x[0]  x[1]  x[2]   x[3]
    y[0]  y[1]
-----------------------------
    1     2     3     4     ← x := []int{1, 2, 3, 4}
    1     2    30     4     ← y = append(y, 30)
  {-------xの範囲---------}  → [1 2 30 4]
  {----yの範囲----}          → [1 2 30]
  y[0]  y[1]  y[2]         
*/
