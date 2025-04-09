package main

import "fmt"

func main() {  //liststart
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "ゴルーチンから送信した文字列"
		fromMain := <-ch2
		fmt.Println("無名関数:", fromMain)
	}()
	
	var fromGoroutine string
	select { // チャネルでのやり取りをselectで囲む
	case ch2 <- "mainから送信した文字列": // こちらは最初は書き込めない
	case fromGoroutine = <-ch1: // ch1に書き込まれればこれが実行される
	}
	fmt.Println("main:", fromGoroutine)
}  //listend
