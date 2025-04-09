package main

import (
	"fmt"
	"sync"
)

func main() {
	// Parseは2度呼ばれているが「初期化中！」は一度しか出力されない
	result := Parse("hello")
	fmt.Println(result)
	result2 := Parse("goodbye")
	fmt.Println(result2)
}

type SlowComplicatedParser interface {
	Parse(string) string
}

// 「func() SlowComplicatedParser」型のパッケージレベルの変数initParserCachedの宣言と初期値の代入  //liststart
var initParserCached func() SlowComplicatedParser =
	sync.OnceValue(initParser) // initParserを1回呼び出す関数が返る（キャッシュされる）


func Parse(dataToParse string) string {
	parser := initParserCached()
	return parser.Parse(dataToParse)
}  //listend

/* 参考 ex1214.goの定義
var parser SlowComplicatedParser // パッケージレベル
...   
func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser() 
	})
	return parser.Parse(dataToParse) // parserを使って「解析」
}
*/


func initParser() SlowComplicatedParser {
	// 初期化処理を行う
	fmt.Println("初期化中!")
	return SCPI{}
}

type SCPI struct { //  Slow Complicated Parser Instance
}

// SCPI のメソッド
func (s SCPI) Parse(in string) string {
	if len(in) > 1 { // 先頭1文字を返す
		return in[0:1]
	}
	return ""
}
