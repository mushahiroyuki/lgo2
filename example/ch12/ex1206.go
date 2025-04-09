package main

import (
	"fmt"
)

func countTo(max int) <-chan int { // intのチャネルを戻す //liststart
	ch := make(chan int)

	go func() { // 無名関数を使ったゴルーチンの呼び出し
		for i := 0; i < max; i++ {
			ch <- i // 0からmax-1までの値を順にchに入れる
		}
		close(ch)
	}()

	return ch // チャネルchを戻す
}

func main() {
	for i := range countTo(10) { // countToからチャネルが戻り、0...9が代入される
		fmt.Println(i)
	}
	fmt.Println("main終了")
}  //listend

/* 次のようにするのと同じ。上のほうが（慣れれば）簡潔
func main() {
	ch := countTo(10) // int のチャネルが戻る
	for i := range ch { // クローズされるまで読み込まれる
		fmt.Print(i, " ")
	}
	fmt.Println()
}
*/
