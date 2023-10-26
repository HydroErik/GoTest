package main

import (
	"fmt"
	"log"
	"net/http"
)

func demoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Demo")
	//fmt.Fprintf(w, "")
}

func main() {

	http.HandleFunc("/demo/", demoHandler)
	fmt.Println("Server Running")

	log.Fatal(http.ListenAndServe(":80", nil))
	//log.Fatal(http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/sites.hydrologik.net/fullchain.pem", "/etc/letsencrypt/live/sites.hydrologik.net/privkey.pem", nil))
}
