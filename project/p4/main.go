package main

import (
    "github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	Age      string
	Sex      string
	Location string
}

func userConstructor(name string, age string, sex string, location string) string {
	user := User{
		Name:     name,
		Age:      age,
		Sex:      sex,
		Location: location,
	}

	encodeToJSon, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error: ", err)
	}

	finalPackage64Encode := base64.URLEncoding.EncodeToString(encodeToJSon)
	return finalPackage64Encode

}

func handleFunc(res http.ResponseWriter, req *http.Request) {
	templatePage, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	name := req.FormValue("name")
	age := req.FormValue("age")
	sex := req.FormValue("sex")
	location := req.FormValue("location")

	encodedData := userConstructor(name, age, sex, location)

	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			// Secure: true,
			Name:     "session-fino",
			Value:    id.String() + "," + name + "," + age + "," + sex + "," + location + "," + encodedData,
			HttpOnly: true,
		}
	}

	http.SetCookie(res, cookie)
	templatePage.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
