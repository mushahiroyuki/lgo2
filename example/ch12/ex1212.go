// ex1212b.goに拡張版があります。
package main

import (
	"fmt"
	"sync"
	"math/rand"	
	"time"	
)

func main() { //liststart1
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		doThing1()
	}()

	go func() {
		defer wg.Done()
		doThing2()
	}()

	go func() {
		defer wg.Done()
		doThing3()
	}()

	wg.Wait()
} //listend1

func doThing1() {
	time.Sleep(getSec(1))
	fmt.Println("doThing1 終了!")
}

func doThing2() {
	time.Sleep(getSec(2))
	fmt.Println("doThing2 終了!")
}

func doThing3() {
	time.Sleep(getSec(3))
	fmt.Println("doThing3 終了!")
}

// スリープする時間を戻す
func getSec(id int) time.Duration {
	sec := rand.Intn(6)+1 // 1以上7未満の整数が返る
	fmt.Printf("doThing%d: %v秒\n", id, sec)
	return time.Duration(time.Duration(sec) * time.Second)
	//  secをtime.Durationに型変換してからtime.Secondeと乗算
}
