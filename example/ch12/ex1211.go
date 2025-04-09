package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 時間制限が課せられる関数（実際の仕事をする関数。ここでは単純な例） //liststart1
func doSomeWork() int {
	x := rand.Intn(20) // 0以上20未満の整数が返る
	fmt.Println("x:", x)
	if x%2 == 0 {  // 2で割り切れればその値を返す
		return x
	} else {  // それ以外の場合は3秒スリープして100を返す
		// 制限時間が2秒なので、オーバーすることになる
		time.Sleep(3 * time.Second) // 3秒スリープ
		return 100
	}
}

// 関数timeLimit -- 引数workerに指定された関数（こではdoSomeWork）を実行し
// 型Tが戻り値の関数workerを実行し、limit以内に型Tの結果を返す
// Tは任意（any）の型に置き換わる「型パラメータ」
// この例ではTがintになって実行される（「8章 ジェネリクス」参照）
func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)  // 値を送るための型Tのチャネル
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	// 実行時間をlimitに制限するcontextを生成
	defer cancel() // 関数を抜ける際に実行。contextのリークを防止

	go func() { // ゴルーチン。select以下とは並行に実行される
		out <- worker() // 引数に指定された（実際の仕事をする）関数を呼び出す
		// 結果はチャネルoutに入れる
	}()

	select {
	case result := <-out:  // チャネルoutを介して結果をもらい、
		return result, nil // 結果とnilを返す
	case <-ctx.Done():  // 制限時間経過した
		var zero T  // zeroは型Tのゼロ値で初期化される
		return zero, errors.New("制限時間オーバー")
	}
}

func main() {
	result, err := timeLimit(doSomeWork, 2*time.Second)
	// doSomeWorkが2秒以内に実行されればresultに値が戻る
	// 制限時間オーバーしてしまえばerrにnil以外の値が入る
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("結果: %v\n", result)
	}
}  //listend1
