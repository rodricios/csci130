/* 
create a web page which serves at localhost over https using TLS
*/

package main

import (
	"fmt"
	"net/http"
)

func servePage(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,
	`<body>
        <h1>TLS</h1>
    </body>`)
}

func main(){
	http.HandleFunc("/", servePage)
	http.ListenAndServeTLS(":8080", "../ocert.pem", ".//key.pem", nil)
}