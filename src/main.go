package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("listening on here")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		resp, err := http.Get("https://search-movies-cu4klqobhcz7knhi47dwylb3yu.eu-west-1.es.amazonaws.com/movies/_search?q=burton")
		if err != nil {
			// handle error
			fmt.Fprintf(w, "handle error")

		}

		//body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		//log.Println(string(body))

		//var result map[string]interface{}
		//json.NewDecoder(resp.Body).Decode(&result)
		//log.Println(result)

		var text, error = ioutil.ReadAll(resp.Body)

		log.Println(string(text))
		log.Println(error)

		//
		defer resp.Body.Close()
		fmt.Fprintf(w, string(text))
	})

	http.ListenAndServe(":3000", nil)
}
