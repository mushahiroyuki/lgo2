package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// タイムリミットが課せられる（実際の仕事をする）関数（ここでは単純な例） //liststart1
func doSomeWork() int {
	x := rand.Intn(20) // 0以上20未満の整数を返す
	fmt.Println("x:", x)
	if x%2 == 0 {  // 2で割り切れればその値を返す
		return x
	} else {  // それ以外の場合は2秒スリープして100を返す
		time.Sleep(2 * time.Second) 
		return 100
	}
}

// timeLimit -- workerに指定された関数を実行しlimit以内に結果を返す
// この例ではworkerとしてdoSomeworkが指定される
// Tはintなど任意（any）の型に置き換わる「型パラメータ」（「8章 ジェネリクス」参照）
func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T)  // 値を送るための型Tのチャネル、キャパシティ1
	ctx, _ := context.WithTimeout(context.Background(), limit)
	// 実行時間をlimitに制限するcontextを生成。

	// defer cancel()
	// じつは関数cancelを明示的に呼び出さなくても、所定の時間が経過すると
	// チャネルctx.Done()から読み込めるようになるため *1*のoutが
	// 準備できる前に *2* のctx.Done()のチャネルが閉じられてしまうため
	// *2* の本体が実行され、エラーが戻されます
	// ただし、このコードだと、go vet を実行すると警告が出力されます
	// contextのリークを防止するためにcancel()を呼び出すほうがお作法にはかなっているようです

	go func() { // ゴルーチン。select以下とは独立に実行される
		out <- worker() // 引数に指定された（実際の仕事をする）関数を呼び出す
		// 結果はチャネルoutに入れる
	}()
	select {
	case result := <-out: // *1*  // チャネルoutから処理結果をもらい、
		return result, nil // 結果を返す。呼び出し先でerrがnilになる
	case <-ctx.Done(): // *2* 
		var zero T  // zeroは型Tのゼロ値で初期化される
		return zero, errors.New("処理時間オーバー")
	}
}

func main() {
	result, err := timeLimit(doSomeWork, 1*time.Second)
	// doSomeWorkが1秒以内に実行されればresultに値が戻る
	// タイムアウトしてしまえばerrにnil以外の値が入る
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("結果: %v\n", result)
	}
}  //listend1
