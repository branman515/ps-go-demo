package main

import (
	"io"
	"net/http"
	"os"
)

// You will need to run this program within the directory of this program
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
