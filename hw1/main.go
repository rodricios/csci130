//Create a template that uses conditional logic
package main

import (
	"log"
	"os"
	"text/template"
)

type payload struct {
	FileName string
	Number int
    Condition bool
}

func main() {
	p1 := payload{
		FileName: "data1",
        Number: 9999,
	}
    
    if p1.Number > 1 {
        p1.Condition = true
    } 

	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
