package main

import (
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DistribuTrace Collector Running")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/spans", receiveSpan)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func receiveSpan(
	w http.ResponseWriter,
	r *http.Request,
) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println(
		"Span received: ",
		string(body),
	)
	w.WriteHeader(http.StatusAccepted)
}