package main

import (
	"fmt"
	"sync"
)

// processAndGather
// T, R: 型パラメータ 「8章 ジェネリクス」参照
// num個のコルーチンを並行に実行する。各ゴルーチンはprocessorを呼んで処理する
func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {  //liststart1
	out := make(chan R, num) //キャパシティnum
	var wg sync.WaitGroup
	// fmt.Printf("wg# 回目  入力  出力\n")
	wg.Add(num) // num個のゴルーチンを並行に実行
	for i := 0; i < num; i++ {

		go func(j int) { // processorを呼んで処理する
			defer wg.Done() // 処理が終わったらwgをデクリメント
			k := 0  // このゴルーチンで何個目の処理か
			for v := range in {
				result := processor(v) // この例の場合、2倍したものをoutに書き込む
				out <- result
				k++
				fmt.Printf("wg#%2v %2v回目 入力:%2v 出力:%2v\n", j, k, v, result)
			}
			fmt.Printf("コルーチン#%vは%v個処理して終了\n", j, k)
		}(i)

	}

	go func() { // モニタリング用のゴルーチン
		wg.Wait() // すべてのゴルーチンが処理を終了するのを待つ
		close(out) // すべてが終了したらチャネルoutをclose
		fmt.Printf("チャネルoutをcloseした\n")
	}()

	var result []R
	for v := range out {
		result = append(result, v)
		fmt.Printf("resultにv（%v）をappendした。result: %v\n", v, result)
	}
	return result
}

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()

	results := processAndGather(ch, func(i int) int {
		return i * 2
	}, 3) // 3は並行実行するゴルーチンの個数
	fmt.Println(results)
}  //listend1
