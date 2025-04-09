package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result, err := timeLimit(doSomeWork, 2*time.Second)
	fmt.Println(result, err)
}

func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {  //liststart1
	out := make(chan T, 1)
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()

	go func() {
		out <- worker()
	}()

	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("work timed out")
	}
}  //listend1


func doSomeWork() int { // 実際の仕事をする関数。ここではダミー
	if x := rand.Int(); x%2 == 0 {  // 2で割り切れればその値を返す
		return x
	} else {  // それ以外の場合は10秒スリープして100を返す
		time.Sleep(10 * time.Second)
		return 100
	}
}
