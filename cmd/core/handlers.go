package core

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go start")
}

func RunHandle() {
	http.HandleFunc("/", home_page)
	fmt.Println("Hello")
	http.ListenAndServe(":8080", nil)
}
