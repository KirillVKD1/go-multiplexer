package main

import (
	"go-multiplexer/fetch"
	"net/http"
)

func main() {
	http.HandleFunc("/", fetch.FetchAll)
	http.ListenAndServe(":8000", nil)
}
