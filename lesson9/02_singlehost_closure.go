package main

import (
	"net/http"
)

func SingleHost(handler http.Handler, allowedHost string) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(ourFunc)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo 2 is Success!"))
}

func main() {
	//single := SingleHost(http.HandlerFunc(myHandler), "webyang.net")
	single := SingleHost(http.HandlerFunc(myHandler), "localhost:8081")

	println("Listening on port 8081")
	http.ListenAndServe(":8081", single)
}
