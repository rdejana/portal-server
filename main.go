package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"portal-server/core"
)

type RequestWrapperImpl struct {
	r *http.Request //The request
}

func (rwi *RequestWrapperImpl) GetPathParameter(param string) string {
	return chi.URLParam(rwi.r, param)
}

func (rwi *RequestWrapperImpl) GetRequest() *http.Request {
	return rwi.r
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	context, handler := Abc()

	r.Get(context, func(w http.ResponseWriter, r *http.Request) {
		wrapper := &RequestWrapperImpl{
			r: r,
		}
		handler(w, wrapper)
	})

	http.ListenAndServe(":3333", r)
}

func FunAndGames(w http.ResponseWriter, req core.RequestWrapper) {
	v := req.GetPathParameter("somevalue")
	w.Write([]byte("hello world with wrapper:" + v))
}

func Abc() (string, func(w http.ResponseWriter, req core.RequestWrapper)) {
	return "/plugins/{somevalue}", FunAndGames

}
