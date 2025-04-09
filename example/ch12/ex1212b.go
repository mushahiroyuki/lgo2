package main

import (
	"fmt"
	"sync"
	"math/rand"	
	"time"	
)

func main() { //liststart1
	var wg sync.WaitGroup
	n := 6; // ゴルーチンの個数
	wg.Add(n)

	for i:=1; i<=n; i++ {
		go func(i int) {
			defer wg.Done()
			doThing(i)
		}(i)
	}

	/*	1.22 以降なら次のコードでも大丈夫だが、上なら以前のバージョンでも大丈夫。「12.5.2 ゴルーチンとforループ」参照
	for i:=1; i<=n; i++ {
		go func() {
			defer wg.Done()
			doThing(i)
		}()
	}
	*/
	
	wg.Wait()
} //listend1

func doThing(i int) {
	start := time.Now() // 処理Aの開始時刻を記録
	time.Sleep(getSec(i))
	fmt.Printf("doThing%d 終了! (%v)\n", i, time.Since(start))
}

// スリープする時間を戻す
func getSec(id int) time.Duration {
	sec := rand.Intn(6)+1 // 1以上7未満の整数が返る
	fmt.Printf("doThing%d: %v秒\n", id, sec)
	return time.Duration(time.Duration(sec) * time.Second)
	//  secをtime.Durationに型変換してからtime.Secondeと乗算
}
