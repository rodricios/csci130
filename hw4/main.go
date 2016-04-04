
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
    
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
	http.HandleFunc("/", surfWebPage)
    
	http.ListenAndServe(":8080", nil)
}

func surfWebPage(res http.ResponseWriter, req *http.Request) {

	var err error

	exersicePage := template.New("index.html")
	exersicePage, err = exersicePage.ParseFiles("index.html")

	if err != nil {
		log.Fatalln(err)
	}

	err = exersicePage.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
	}
}