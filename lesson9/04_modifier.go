package main

import (
	"net/http"
	"net/http/httptest"
)

type ModifierMiddleware struct {
	handler http.Handler
}

func (m *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	// passing a ResponseRecorder instead of the original RW
	m.handler.ServeHTTP(rec, r)
	// after this finishes, we have the response recorded
	// and can modify it before copying it to the original RW

	// we copy the original headers first
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	// and set an additional one
	w.Header().Set("my-blog", "webyang.net")
	// only then the status code, as this call writes the headers as well
	w.WriteHeader(418)
	// the body hasn't been written yet, so we can prepend some data.
	w.Write([]byte("Middleware says hello again. \r\n"))
	// then write out the original body
	w.Write(rec.Body.Bytes())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo4 is Success!"))
}

func main() {
	mid := &ModifierMiddleware{http.HandlerFunc(myHandler)}

	println("Listening on port 8081")
	http.ListenAndServe(":8081", mid)
}
