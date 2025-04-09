package main

import (
	"fmt"
)

func main() {  //liststart
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "ゴルーチンから送信した文字列" // ch1に書き込めない限りここで待たされる
		v1 := <-ch2
		fmt.Println(v1)
	}()

	ch2 <- "mainから送信した文字列"  // ch2に書き込めない限りここで待たされる
	v2 := <-ch1
	fmt.Println(v2)
}  //listend
