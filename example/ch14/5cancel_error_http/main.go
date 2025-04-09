package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func main() {
	ctx, cancelFunc := context.WithCancelCause(context.Background())  //liststart1
	defer cancelFunc(nil) //listend1
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {  
		defer wg.Done()
		for {
			resp, err := makeRequest(ctx, "http://httpbin.org/status/200,200,200,500")  //liststart2
			if err != nil {
				cancelFunc(fmt.Errorf("in status goroutine: %w", err))
				return
			}
			if resp.StatusCode == http.StatusInternalServerError {
				cancelFunc(errors.New("bad status"))
				return
			}
			ch <- "success from status"
			time.Sleep(1 * time.Second)  //listend2
		}
	}()
	
	go func() {
		defer wg.Done()
		for {
			resp, err := makeRequest(ctx, "http://httpbin.org/delay/1")  //liststart3
			if err != nil {
				fmt.Println("in delay goroutine:", err)
				cancelFunc(fmt.Errorf("in delay goroutine: %w", err))
				return
			}
			ch <- "success from delay: " + resp.Header.Get("date") //listend3
		}
	}() 
	
loop:  //liststart4
	for {
		select {
		case s := <-ch:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			fmt.Println("in main: cancelled with error", context.Cause(ctx))
			break loop
		}
	}
	wg.Wait()
	fmt.Println("context cause:", context.Cause(ctx))  //listend4
}
