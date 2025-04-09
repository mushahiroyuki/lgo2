package main

import (
	"fmt"
	"sync"
)

func main() {
	// "初期化中！" は一度だけ表示される
	result := Parse("hello")
	fmt.Println(result)
	result2 := Parse("goodbye")
	fmt.Println(result2)
}

//「遅くて複雑な」構文解析器 //liststart1
type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser // パッケージレベル（関数の外）
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser() // 一度だけ行われる
	})
	return parser.Parse(dataToParse) // parserを使って「解析」
}

func initParser() SlowComplicatedParser {
	// 諸々のセットアップやロードをここで行う
	fmt.Println("初期化中！")
	return SCPI{}
}

type SCPI struct {
}

// 型SCPIのメソッドParse（構文解析を行う）
func (s SCPI) Parse(in string) string {
	if len(in) > 1 { // ここでは単に最初の1文字を返す
		return in[0:1]
	}
	return ""  // 空文字列のときは空文字列を返す
}  //listend1
