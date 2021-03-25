
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

// This represents a simple example of how Go can be used to interact with the Bitly API 

// Here we define a struct that represents the fields that come back from the GET Bitlink request in the Bitly API
// API Docs: https://dev.bitly.com/api-reference#getBitlink
type Bitlink struct {
	Link string `json:"link"`
	CreatedBy string `json:"created_by"`
	LongUrl string `json:"long_url"`
}

func main() {
	// set up http client to make the http requests to the api
	client := &http.Client{}

	// create a new HTTP request object
	req, err := http.NewRequest("GET", "https://api-ssl.bitly.com/v4/bitlinks/bit.ly/BITLINK_ID", nil)
	if err != nil {
		log.Fatal(err)
	}

	// set the authorization header - this is how you prove to the Bitly API that you are who you
	// say you are - think of it as a password of sorts
	req.Header.Set("Authorization", "Bearer AUTH_TOKEN")

	// Make the HTTP request to the Bitly endpoint
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read the body of the request into a [] of bytes
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// We take the now []bytes that represents the body of the request we made to the API 
	// and we map it into a bitlink struct which we can then view
	var link Bitlink
	err = json.Unmarshal(bodyText, &link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", link)
}















