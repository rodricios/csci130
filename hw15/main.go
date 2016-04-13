/*
Create a webpage which writes a cookie to the client's machine. 

This cookie should be designed to create a session and should use a UUID, HttpOnly, and Secure (though you'll need to comment secure out).
*/
package main

import (
  "net/http"
  "crypto/hmac"
  "crypto/sha256"
  "io"
  "fmt"
)

var SESSION_DIGEST string = ""

const SESSION_DATA = "DADA"

const SESSION_COOKIE = "SESSION"

const PRIVATE_KEY = "PRIVATE_KEY"

func index(responseWriter http.ResponseWriter, request *http.Request) {
    sessionCookie, cookieError := request.Cookie(SESSION_COOKIE)
  
	if cookieError == http.ErrNoCookie {
        //uid, _ := uuid.NewV4()
        sessionCookie = &http.Cookie{
        Name: SESSION_COOKIE,
        Value: SESSION_DATA,
        HttpOnly: true,
        //Secure: true
        }
        // Encrypt data
        SESSION_DIGEST = encrypt(SESSION_DATA)
    } else {
        if encrypt(sessionCookie.Value) == SESSION_DIGEST {
            fmt.Println("Cookie was not changed!")
        } else {
            fmt.Println("Cookie was changed!")
        }
    }
    http.SetCookie(responseWriter, sessionCookie)
}

func encrypt(data string) string {
	h := hmac.New(sha256.New, []byte(PRIVATE_KEY))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}