package main

import "fmt"

type Direction int  //liststart1

const (
	_ Direction = iota
	North
	South
	East
	West
)

//go:generate stringer -type=Direction

func main() {
	fmt.Println(North.String())
}  //listend1
