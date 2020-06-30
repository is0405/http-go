package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"net/url"
	"strings"
)

type json_data struct {
	IP string `json:"IP"`
}

//var sc = bufio.NewScanner( os.Stdin )

func main() {

	//var url_h = connect_url()
	var url_h = "http://localhost:8080"
	url_h += "/test"

	val:= url.Values{}
	val.Set("key", "value")
	req, err := http.NewRequest("POST",url_h, strings.NewReader(val.Encode()))
	if err != nil {
		fmt.Println( "Error:http" )
		fmt.Println( err )
		os.Exit( 0 )
	}
	req.Header.Set("Authorization", "XXXXXXXX")

	client := new(http.Client)
  resp, err := client.Do(req)
	if err != nil {
		fmt.Println( "Error:resp" )
		fmt.Println( err )
		os.Exit( 0 )
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println( "Error:resp" )
		fmt.Println( err )
		os.Exit( 0 )
	}
	fmt.Println("[body] " + string(body))

}

func connect_url() string {
	var url = "http://"
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println( "json_file read Error" )
		fmt.Println( err )
		os.Exit( 1 )
	}

	var data json_data

	json.Unmarshal( raw, &data )
	url += data.IP

	return url
}
/*
func input() string {
    sc.Scan()
    return sc.Text()
}*/
