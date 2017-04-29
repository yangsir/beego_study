package main

import (
	"net/http"
)

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

/*func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}*/

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo 1 is Success!"))
}

func main() {
	//single := NewSingleHost(http.HandlerFunc(myHandler), "localhost:8081")

	single := &SingleHost{
		handler:     http.HandlerFunc(myHandler),
		allowedHost: "webyang.net",
	}

	println("Listening on port 8081")
	http.ListenAndServe(":8081", single)
}
