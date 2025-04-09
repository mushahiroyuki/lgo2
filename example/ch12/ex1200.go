package main

import "fmt"

func doBusinessLogic(val int) int {  //liststart1
	return val*val // この場合は単純に「自乗する」のがビジネスロジック
}

func runThingsConcurrently(chIn <-chan int, chOut chan<- string) { 
	for val := range chIn { // チャネルchInから値が到着するたびに
		go func(val int) {  // 新たにゴルーチンを呼び出す（無名関数）
			result := doBusinessLogic(val) // valを処理して結果をresultに代入
			chOut <- fmt.Sprintf("%d -> %d", val, result)
			// 結果の文字列の生成してチャネルchOutに書き込み
		}(val)  // chInから読んだ値を無名関数の引数として渡して実行
	}
}

func main() {
	chIn := make(chan int)  // intのチャネルを生成。送信用
	chOut := make(chan string) // stringのチャネルを生成。受信用
	go runThingsConcurrently(chIn, chOut)

	vals := []int{1, 2, 3, 4, 5}
	for v := range vals { // スライスvalsの各要素に対して
		chIn <- v   // その値をchInに送信
	}

	i := 0
	for v := range chOut { // chOutから受信するごとに
		fmt.Println(v) // 結果の文字列を出力。順番は不定
		i++
		if len(vals) <= i { // 送った要素の数と同じだけ到着したら
			break  // forループを抜ける
		}
	}

	close(chOut) // chOutをクローズ（終了してしまうので、なくても大丈夫ではある）
	close(chIn) // chOutをクローズ（同上）
}  //listend1
