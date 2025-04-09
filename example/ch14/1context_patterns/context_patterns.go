package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// コンテキストの生成方法と、関数への渡し方を示す
func main() {
	ctx := context.Background() // コンテキストの生成
	result, err := logic(ctx, "a string")
	fmt.Println(result, err)
}

// コンテキストをパスする（あるいは使う）関数用の引数を示す
func logic(ctx context.Context, info string) (string, error) {
	// ここで処理
	
	return "", nil
}

// Middleware -- コンテキストの利用法を示す例
func Middleware(handler http.Handler) http.Handler {  //liststart
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()  // ① コンテキストを抽出 

		// 【 ② ここでctxに情報を追加（詳細はこのあと）】

		req = req.WithContext(ctx)  // ③ コンテキストを伴った新しいリクエストを生成
		handler.ServeHTTP(rw, req) // ④ reqにコンテキストが含まれる
	})
}  //listend

// *http.Requestからコンテキストを抽出し、関数に渡す
func handler(rw http.ResponseWriter, req *http.Request) {  //liststart2
	ctx := req.Context()
	err := req.ParseForm() // req.Formとreq.PostFormに入れる
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data) // 必要な処理を行う
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}  //listend2

type ServiceCaller struct {  //liststart3
	client *http.Client
}

// *http.Requestにコンテキストを追加
func (sc ServiceCaller) callAnotherService(
			ctx context.Context, data string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d",
			resp.StatusCode)
	}
	// レスポンスを処理
	id, err := processResponse(resp.Body)
	return id, err
}  //listend3

// レスポンスを処理する（仮のコード）
func processResponse(body io.ReadCloser) (string, error) {
	return "", nil
}
