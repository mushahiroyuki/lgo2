/*
ex1210x.go -- ex1210.goの詳しい説明
別のターミナルでgo run ex1210.go を実行しておいて   
ex1210.sh を実行すると12個のリクエストを送ります。
*/

package main

import (
	"errors"
	"net/http"
	"time"
	"fmt"
)

// 空（くう）の構造体（empty struct。フィールドがない構造体）。
// 空の構造体はメモリを消費しない
type EmptyStruct struct{}

// 構造体の型PressureGaugeを定義（pressure gaugeは「圧力計」）
type PressureGauge struct { // 「圧力計」  //liststart1
	ch chan EmptyStruct
	// 型が「空の構造体」のデータを保持するチャネル。フィールド名はch
	// 空の構造体はメモリを消費しないので、「シグナル」として用いる
}

// PressureGaugeのメソッド。サイズlimitのpressure gaugeを戻す
func New(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan EmptyStruct, limit),
		// フィールドchを初期化。バッファ付きのチャネル
	}
}

// PressureGauge pgを受け取って、それが許す範囲で関数fを実行する
func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- EmptyStruct{}: // チャネルpg.chに書き込み
		f()
		<-pg.ch  // pg.chから値をもらったが値は無視（シグナルとして受信）
		return nil
	default:
		return errors.New("キャパシティに余裕がありません")
	}
}

// 実行される個数が制限される作業
// ここでは単に「2秒経ったら"完了" を返す」
func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "完了\n"
}

func main() {
	pg := New(5)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
			// []byteへの型変換。
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
