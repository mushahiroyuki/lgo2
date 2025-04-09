package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/learning-go-book-2e/ch14/sample_code/context_guid/tracker"
	"io"
	"net/http"
)

type Logic interface {
	Process(ctx context.Context, data string) (string, error)
}
type Controller struct {
	Logic Logic
}

func (c Controller) First(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	data := req.URL.Query().Get("data")
	result, err := c.Logic.Process(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

type Logger interface {  //liststart3
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type LogicImpl struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}  //listend3

func (l LogicImpl) Process(ctx context.Context, data string) (string, error) {  //liststart4
	l.Logger.Log(ctx, "starting Process with "+data)
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, l.Remote+"/second?query="+data, nil)
	if err != nil {
		l.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	req = l.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)  //listend4
	if err != nil {
		l.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	if resp.Body == nil {
		l.Logger.Log(ctx, "empty response from second")
		return "", nil
	}
	defer resp.Body.Close()
	out, err := io.ReadAll(resp.Body)
	return string(out), err
}

func main() {
	r := chi.NewRouter()
	r.Use(tracker.Middleware)
	controller := Controller{  //liststart1
		Logic: LogicImpl{
			RequestDecorator: tracker.Request,
			Logger:           tracker.Logger{},
			Remote:           "http://localhost:4000",
		},
	}  //listend1
	r.Get("/first", controller.First)
	http.ListenAndServe(":3000", r)
}
