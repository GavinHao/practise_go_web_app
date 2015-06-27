package main

import (
	"net/http"
	"log"
)

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/",someFunc)
	http.ListenAndServe(":8800",myMux)
}

func someFunc(w http.ResponseWriter,r *http.Request) {
	path := r.URL.Path[0:]
	log.Println(path)
	w.Write([]byte("hello"+r.URL.Path[0:]))
}
