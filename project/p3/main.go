/*
PROJECT STEP 3 - continuing to build our application, create a template 
which is a form. 
The form should gather the user's name and age. 
Store the user's name and age in the cookie.
*/

package main

import (
  "github.com/nu7hatch/gouuid"
  "net/http"
)

const SESSION_COOKIE = "session-fino"

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