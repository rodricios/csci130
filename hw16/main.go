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

func redirectTLS(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, "https://127.0.0.1:10443/"+req.RequestURI, http.StatusMovedPermanently)
}

func main(){
	http.HandleFunc("/", servePage)
	//key and cert Im not going to upload to Github
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}