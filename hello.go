package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
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
