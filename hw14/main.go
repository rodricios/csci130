/*
Create a webpage which writes a cookie to the client's machine. 

This cookie should be designed to create a session and should use a UUID, HttpOnly, and Secure (though you'll need to comment secure out).
*/
package main

import (
  "github.com/nu7hatch/gouuid"
  "net/http"
)

const SESSION_COOKIE = "SESSION"

func index(responseWriter http.ResponseWriter, request *http.Request) {
	sessionCookie, cookieError := request.Cookie(SESSION_COOKIE)
  
	if cookieError == http.ErrNoCookie {
    uid, _ := uuid.NewV4()
    sessionCookie = &http.Cookie{
      Name: SESSION_COOKIE,
      Value: uid.String(),
      HttpOnly: true,
      //Secure: true
    }
  }
  
	http.SetCookie(responseWriter, sessionCookie)
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}