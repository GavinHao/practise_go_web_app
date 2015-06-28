package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
)

type Myhandler struct  {
}

func (this *Myhandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	paths := r.URL.Path[1:]
	log.Println(paths)

	var filePath string
	if strings.HasPrefix(paths,"public"){
		filePath = paths
	}else{
		filePath = "template/"+paths
	}
	log.Println(filePath)
	data,err:=ioutil.ReadFile(filePath)

	if err == nil {
		var contentType string
		if strings.HasSuffix(paths,".css"){
			contentType = "text/css"
		}else if strings.HasSuffix(paths,".js"){
			contentType = "application/javascript"
		}else if strings.HasSuffix(paths,".png"){
			contentType = "image/png"
		}else if strings.HasSuffix(paths,".html"){
			contentType = "text/html"
		/*}else if strings.HasSuffix(paths,""){
			contentType = "text/css"
		}else if strings.HasSuffix(paths,".css"){*/

		}else{
			contentType = "text/plain"
		}

		w.Header().Add("content type",contentType)
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
