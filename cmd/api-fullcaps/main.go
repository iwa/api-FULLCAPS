package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", routeRoot)
	http.HandleFunc("/uppercase", routeUppercase)
	http.HandleFunc("/lowercase", routeLowercase)

	fmt.Println("Server started at localhost:8080")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func routeRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		fmt.Fprint(w, "why delete me? :(")
		return
	}
	fmt.Fprint(w, "hello :)")
}

func routeUppercase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	bodyString := string(bodyBytes)
	bodyStringUpper := strings.ToUpper(bodyString)

	fmt.Fprint(w, bodyStringUpper)
}

func routeLowercase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	bodyString := string(bodyBytes)
	bodyStringUpper := strings.ToLower(bodyString)

	fmt.Fprint(w, bodyStringUpper)
}
