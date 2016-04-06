/*
Create a webpage that serves a form and allows the user to upload a txt file. You do not need to check if the file is a txt; bad programming but just trust the user to follow the instructions. Once a user has uploaded a txt file, copy the text from the file and display it on the webpage. Use req.FormFile and io.Copy to do this
*/
package main
import (
	"net/http"
    "log"
	"io/ioutil"
	"io"
	"os"
)

func main(){
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
        //formPage, _ := template.ParseFiles("fileUpload.html")
        //formPage.Execute(responseWriter,nil)
        responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
        io.WriteString(responseWriter, 
            `<form action ="/file" method="post" enctype="multipart/form-data">
                File:<br>
                <input type="file" name="filename"><br>
                <input type="submit" value="Submit"><br>
            </form>`)
    })
    
	http.HandleFunc("/file", func(responseWriter http.ResponseWriter, request *http.Request) {
        inputFile, fileHeader, error := request.FormFile("filename") 
        log.Println(fileHeader.Filename)
        
        fileCopy, error := os.Create(fileHeader.Filename)
        if error != nil {panic(error)}
        
        io.Copy(fileCopy, inputFile)
        if error != nil {panic(error)}
        
        fileBytes, error := ioutil.ReadFile(fileCopy.Name())
        if error != nil {panic(error)}
        
        fileString := string(fileBytes)
        
        io.WriteString(responseWriter, fileString)
    })
    
	http.ListenAndServe(":8080", nil)
}