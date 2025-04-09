// ex1213x.goでは、途中の様子がわかるようにメッセージを出力しています
package main

import (
	"fmt"
	"sync"
)

// processAndGather
// T, R: 型パラメータ （「8章 ジェネリクス」参照）
// num個のゴルーチンを並行実行する。各ゴルーチンはprocessorを呼んで処理する
func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {  //liststart1
	out := make(chan R, num) //キャパシティnum
	var wg sync.WaitGroup
	wg.Add(num) // num個のゴルーチンを並行に実行
	for i := 0; i < num; i++ {
		go func() { // processorを呼んで処理する
			defer wg.Done() // 処理が終わったらwgをデクリメント
			for v := range in {  // チャネルinから読み込み
				out <- processor(v) // processorの結果をoutに書き込む
			}
		}()
	}

	go func() { // モニタリング用のゴルーチン
		wg.Wait() // すべてのゴルーチンが処理を終了するのを待つ
		close(out) // すべてが終了したらチャネルoutをclose
	}()

	var result []R
	for v := range out {  // outに値が来たら
		result = append(result, v)  // resultの最後に追加
	}
	return result  // 型Rのスライスを戻す
}

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i // 0から19までをチャネルに入れる
		}
		close(ch)
	}()

	results := processAndGather(ch, func(i int) int {
		return i * 2
	}, 3)  // 3個のゴルーチンを並行実行する
	fmt.Println(results)
}  //listend1
