package main

import (
	"fmt"
)

func countTo(max int) <-chan int { //intのチャネルを戻す //liststart
	ch := make(chan int)
	go func() { // 無名関数
		for i := 0; i < max; i++ {
			ch <- i // 0からmax-1までの値を順にchに入れる
		}
		close(ch)
	}()
	return ch // チャネルchを戻す
}

// ex1206.goのほうが簡潔だが、慣れるまではこちらのほうがわかりやすいかと 
func main() {
	ch := countTo(10) // int のチャネルが戻る
	for i := range ch { // クローズされるまで読み込まれる
		fmt.Println(i)
	}
}
