/*
Create a webpage which uses a cookie to track the number of visits of a user.

Display the number of visits. Make sure that the favicon.ico requests are not also incrementing the number of visits.
*/

package main

import (
    "net/http"
	"strconv"
	"io"
    "log"
    "strings"
)

const VISIT_COOKIE = "NUMBER_OF_VISITS"

func visit(responseWriter http.ResponseWriter, request *http.Request) {
    
	visitCookie, cookieError := request.Cookie(VISIT_COOKIE)
	log.Println(request)
    
    if cookieError == http.ErrNoCookie {
		visitCookie = &http.Cookie{Name: VISIT_COOKIE, Value: "0"}
	} else {
        if !strings.Contains(request.URL.Path, "favicon.ico") {
            numVisits, _ := strconv.Atoi(visitCookie.Value)
            visitCookie.Value = strconv.Itoa(numVisits + 1)
        }
	}
	http.SetCookie(responseWriter, visitCookie)
	io.WriteString(responseWriter, visitCookie.Value)
}

func main(){
	http.HandleFunc("/", visit)
	http.ListenAndServe(":8080", nil)
}