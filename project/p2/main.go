/*
PROJECT STEP 2 - have the application write a cookie called
"session-fino" with a UUID.

The cookie should serve HttpOnly and you should have the "Secure" 
flag set also though comment the "Secure" flag out as we're
not using https.
*/

package main

import (
  "github.com/nu7hatch/gouuid"
  "net/http"
)

const SESSION_COOKIE = "session-fino"

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