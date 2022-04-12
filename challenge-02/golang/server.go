package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", FileHandler)

	fmt.Println("Server started, running 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(data))
}
