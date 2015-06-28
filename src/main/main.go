package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

type Myhandler struct  {
}

func (this *Myhandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	paths := r.URL.Path[0:]
	log.Println(paths)

	data,err:=ioutil.ReadFile("template"+paths)
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404-Not Found!"+http.StatusText(404)))
	}

}

func main() {

	http.Handle("/",new(Myhandler))
	http.ListenAndServe(":8800",nil)
}
