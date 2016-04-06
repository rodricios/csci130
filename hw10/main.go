package main
import (
	"net/http"
	"io"
)

func main(){
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		key := "n"
		val := req.FormValue(key)
		io.WriteString(res, val)
	})

	http.ListenAndServe(":8080",nil)
}