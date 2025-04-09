/*
   go run ex1210.go を実行しておいてから、
   別のターミナルで "sh ex1210.sh" を実行すると
   10個のリクエストを送ります。
*/

package main

import (
	"errors"
	"net/http"
	"time"  // time.Sleepを使う
	"fmt"
)

// 構造体の型PressureGaugeを定義（pressure gaugeは「圧力計」）
type PressureGauge struct { // 「圧力計」  //liststart1
	ch chan struct{} // 構造体のフィールドch
	// 型が「空の構造体」のデータを保持するチャネル
	// 空の構造体はメモリを消費しないので、「シグナル」として用いる
}

// PressureGaugeのファクトリ関数
// サイズlimitのpressure gaugeを戻す
func New(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan struct{}, limit),
		// フィールドchを初期化。バッファ付きのチャネル
		// struct{}型（空の構造体）
	}
}

// PressureGauge のメソッド
// *PressureGauge pgを受け取って、それが許す範囲で関数fを実行する
func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}: // チャネルpg.chに書き込み
		f() // Processに引数にとして渡された関数を実行
		<-pg.ch  // pg.chから値をもらったが値は無視（シグナルとして受信）
		return nil
	default:
		return errors.New("キャパシティに余裕がありません")
	}
}

// 実行される個数が制限される作業
// 単に「2秒経ったら、"完了\n" を返す」という作業をする
func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)  // 2秒スリープ
	return "完了\n"
}

func main() {
	pg := New(5)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
			// stringを[]byteへ型変換
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("リクエストが多すぎてさばききれません\n"))
		}
	})
	
	fmt.Println("ブラウザで次を開いてください: 'http://localhost:8080/request'")
	fmt.Println("あるいは 'sh ex1210.sh' を実行してみてください")	
	http.ListenAndServe(":8080", nil)
}  //listend1
