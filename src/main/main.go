package main

import "net/http"

type Person struct  {
	fname string
}

func (this *Person) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("hello:"+this.fname))
}

func main() {
	p := &Person{fname:"gavin"}
	http.ListenAndServe(":8800",p)
}
