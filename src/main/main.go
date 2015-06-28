package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"html/template"
	"log"
)

type Myhandler struct  {
}

func (this *Myhandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	paths := r.URL.Path[1:]
	log.Println(paths)

	var filePath string
	var contentType string

	if strings.HasPrefix(paths,"public"){ //静态资源
		filePath = paths
		data,err:=ioutil.ReadFile(filePath)
		if err == nil {
			if strings.HasSuffix(paths,".css"){
				contentType = "text/css"
			}else if strings.HasSuffix(paths,".js"){
				contentType = "application/javascript"
			}else if strings.HasSuffix(paths,".png"){
				contentType = "image/png"
			}
			w.Header().Add("content type",contentType)
			w.Write(data)
		} else {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
		}
	}else{
		filePath = "template/"+paths
		if strings.HasSuffix(paths,".html"){
			contentType = "text/html"
			tpl,tplErr := template.ParseFiles(filePath)
			if tplErr==nil {
				tpl.Execute(w,paths)
			} else {
				tpl,tplErr = template.ParseFiles("template/err/404.html")
				if tplErr==nil {
					tpl.Execute(w,nil)
				} else {
					log.Println(tplErr)
				}
			}
		} else if strings.Index(paths,".") > 0 {
			contentType = "text/html"
			tpl,tplErr := template.ParseFiles("template/err/404.html")
			if tplErr==nil {
				tpl.Execute(w,nil)
			} else {
				log.Println(tplErr)
			}
		} else{
			contentType = "text/html"
			tpl,tplErr := template.ParseFiles(filePath+".html")
			if tplErr==nil {
				tpl.Execute(w,paths)
			} else {
				tpl,tplErr = template.ParseFiles("template/err/404.html")
				if tplErr==nil {
					tpl.Execute(w,nil)
				} else {
					log.Println(tplErr)
				}
			}
		}
		w.Header().Add("content type",contentType)
	}
}

func main() {
	http.Handle("/",new(Myhandler))
	http.ListenAndServe(":8800",nil)
}
