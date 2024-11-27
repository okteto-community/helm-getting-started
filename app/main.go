package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting hello-world server...")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world from your friends at Okteto!")
}
