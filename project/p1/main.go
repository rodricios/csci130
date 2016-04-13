/*
PROJECT STEP 1 - create a web application that serves an HTML template.
*/
package main
import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}
    
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err = tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
