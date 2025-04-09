package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	fmt.Println("url:", url)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background()) // キャンセル付きコンテキスト生成 //liststart1
	defer cancelFunc() // 確実にcancel()呼び出すためにdeferを使う //listend1
	ch := make(chan string)  //liststart2
	var wg sync.WaitGroup
	wg.Add(2)  //listend2

	go func() {  //liststart3
		defer wg.Done()
		for {
			resp, err := makeRequest(ctx, "http://httpbin.org/status/200,200,200,500")
			if err != nil {
				fmt.Println("error in status goroutine:", err)
				cancelFunc()
				return
			}
			if resp.StatusCode == http.StatusInternalServerError {
				fmt.Println("bad status, exiting")
				cancelFunc()
				return
			}
			select {
			case ch <- "success from status goroutine":
			case <-ctx.Done():
			}
			time.Sleep(1 * time.Second)
		}
	}()  //listend3

	go func() {  //liststart4
		defer wg.Done()
		for {
			resp, err := makeRequest(ctx, "http://httpbin.org/delay/1") // 1秒後に戻る
			if err != nil {
				fmt.Println("error in delay goroutine:", err)
				cancelFunc()
				return
			}
			select {
			case ch <- "success from delay: " + resp.Header.Get("date"):
			case <-ctx.Done():
			}
		}
	}()  //listend4

loop:  //liststart5
	for {
		select {
		case s := <-ch:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			fmt.Println("in main: cancelled!")
			break loop
		}
	}
	wg.Wait()  //listend5
}
