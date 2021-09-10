package main


import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home_page)
	fmt.Println("Hello")
	http.ListenAndServe(":8080", nil)

}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go start")
}
