package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"log"
	"net/http"
)

var template1 *template.Template

type User struct {
	Uuid, Username, password, Hmac string
	Age                            int
	Valid, LoggedIn                bool
}

func init() {
	var err error
	template1, err = template.ParseFiles("index.html")
	if err != nil {
		log.Println("Error: ", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	cookie, err1 := req.Cookie("user-info")
	user := User{}
	if err1 == http.ErrNoCookie {
		uuid, _ := uuid.NewV4()
		user = User{
			Uuid: uuid.String(),
			Hmac: getCode(uuid.String()),
		}
		log.Println("UUID: ", user.Uuid)
		log.Println("HMAC: ", user.Hmac)
		encodedUser := encodeJsonData(user)
		log.Println("ENCODED: ", encodedUser)
		cookie = setNewCookie(encodedUser, cookie)
		http.SetCookie(res, cookie)
		user.Valid = true
	}
	if req.Method == "POST" {
		log.Println("POST REQUEST")
		Username := req.FormValue("Username")
		password := req.FormValue("password")

		user = User{
			Username: Username,
			password: password,
		}
		var err error
		cookie, err = req.Cookie("user-info")
		if err != nil {
			log.Println("ERROR: ", err)
		}
		cookie.Value = updateCookie(Username, req, cookie)
		user = decodeJsonData(cookie)
	}
	log.Println("Before template data looks like: ", user)
	template1.Execute(res, user)
}

func setNewCookie(userInfo string, cookie *http.Cookie) *http.Cookie {
	cookie = &http.Cookie{
		Name:     "user-info",
		Value:    userInfo,
		HttpOnly: true,
		//		Secure: true,
	}
	return cookie
}

func updateCookie(user User, req *http.Request, cookie *http.Cookie) string {
	decodedUser := decodeJsonData(cookie)
	if decodedUser.Valid == false {
		return encodeJsonData(user)
	}
	user.Uuid = decodedUser.Uuid
	user.Hmac = getCode(user.Uuid + user.Username)
	return encodeJsonData(user)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func encodeJsonData(user User) string {
	jsonUser, errJsonMarshalError := json.Marshal(user)
	if errJsonMarshalError != nil {
		log.Println("Error: ", errJsonMarshalError)
	}
	return base64.StdEncoding.EncodeToString(jsonUser)
}

func decodeJsonData(cookie *http.Cookie) User {
	log.Println("Cookie", cookie.Value)
	decode, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Println("Error: ", err)
		var user User
		user.Valid = false
		return user
	}
	var user User
	json.Unmarshal(decode, &user)
	if user.Hmac == getCode(user.Uuid+user.Username) {
		log.Println("")
		user.Valid = true
		return user
	}
	log.Println("auth fails")
	user.Valid = false
	return user
}
