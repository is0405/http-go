package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	for k, v := range r.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// POST (form)
	if method == "POST" {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println( "Error:http" )
			fmt.Println( err )
			os.Exit( 0 )
		}

		fmt.Println("[request body row] " + string(body))
		decoded, error := url.QueryUnescape(string(body))
		if error != nil {
			fmt.Println( "Error:http" )
			fmt.Println( err )
			os.Exit( 0 )
		}
		fmt.Println("[request body decoded] ", decoded)
		fmt.Fprint(w, decoded,":Recieved Post(form) request!!")
	}
}

func main() {
	http.HandleFunc("/test", handler)
	fmt.Println("webserver start.")
	err := http.ListenAndServe("localhost:8080", nil) //test
	//err := http.ListenAndServe(":80", nil) //AWS
	if err != nil{
		fmt.Println( err )
	}
}
