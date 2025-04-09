package main

import (
	"context"
	"fmt"
	"time"
)

func countTo(ctx context.Context, max int) <-chan int {  //liststart1
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done(): // ctx.Done()が返すチャネルから読み込めた場合（★1）
				fmt.Println("ゴルーチンを抜ける")
				return // ゴルーチンを抜ける
			case ch <- i: // iをチャネルchに書き込む
				fmt.Printf("chに%dを書き込み\n", i)
			}
		}
	}()
	
	return ch // chを戻してcountToは終了。途中で呼び出したゴルーチンは実行を続ける
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// コンテキストを初期化。コンテキストとキャンセルするための関数を返す
	defer cancel()  // 処理が終わったらキャンセルを呼び出す
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break // breakすると、deferされていたcancel()が実行される
			// これによりctx.Done()のチャネルが閉じられ、上の★1のcaseが選択される
		}
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main終了")
}  //listend1
