package main

import (
	"io"
	"log" // Log out things to console
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":8080", nil)
}

func indexHandle(res http.ResponseWriter, req *http.Request) {
    key := "name-input"
    val := req.FormValue(key)
    log.Println(val)
    res.Header().Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(res, 
    `<form method="POST">
        <label for="Name">Enter your name here: </label>
        <input type="text" name="name-input">
        <input type="submit" value="Click here to display your name!">
    </form>`)
    if val != "" {
        io.WriteString(res, "Your name is: "+val)
    }
}