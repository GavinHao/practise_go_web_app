package main

import "net/http"

func main() {
	http.HandleFunc("/",someFunc)
	http.ListenAndServe("localhost:8800",nil)
}

func someFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
