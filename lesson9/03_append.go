package main

import (
	"net/http"
)

type AppendMiddleware struct {
	handler http.Handler
}

func (a *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Middleware says hello! \r\n"))
	a.handler.ServeHTTP(w, r)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo 3 is Success!"))
}

func main() {
	mid := &AppendMiddleware{http.HandlerFunc(myHandler)}

	println("Listening on port 8081")
	http.ListenAndServe(":8081", mid)
}
