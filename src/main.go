package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("listening on")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		resp, err := http.Get("https://httpbin.org/get")
		if err != nil {
			// handle error
			fmt.Fprintf(w, "handle error")

		}

		//body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		//log.Println(string(body))

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		log.Println(result)
		//
		defer resp.Body.Close()
		fmt.Fprintf(w, "Hello World! I'm a HTTP server!")
	})

	http.ListenAndServe(":3001", nil)
}
