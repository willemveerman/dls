package main

import ("net/http"
	"fmt"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}


func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//hello2 := HandlerFunc(hello)

	http.HandleFunc("/hello", hello)

	server.ListenAndServe()

}

