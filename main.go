package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RequestBody struct {
	Commands []Command
}

type Command struct {
	Verb string `json:"verb"`
	Url  string `json:"url"`
	Body struct {
		field1 string
	}
}

const maxCommandsAmount int = 2

func multiplexer(w http.ResponseWriter, req *http.Request) {
	var parsedBody *RequestBody
	err := json.NewDecoder(req.Body).Decode(&parsedBody)
	fmt.Println("parsedReq@@@@@", parsedBody)

	commandsLength := len(parsedBody.Commands)

	if commandsLength > maxCommandsAmount {
		fmt.Println("@@@BAD_REQUEST", err)
		http.Error(w, "max commands amount should be less than"+" "+strconv.Itoa(maxCommandsAmount), http.StatusBadRequest)
		return
	}

	for i := 0; i < commandsLength; i++ {
		fmt.Println("@@@", i, parsedBody.Commands[i])
	}
	fmt.Println("@@@ERROR", err)

	w.Write([]byte("success"))
}

func send() {

}

func main() {
	http.HandleFunc("/", multiplexer)
	http.ListenAndServe(":8091", nil)
}

/* package main

import (
	"fmt"
	"net/http"
)

type home struct{}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("@@@@", r.GetBody())
	w.Write([]byte("Hello, World!"))
}

type page struct {
	body string
}

func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// echo back the page first URI
	fmt.Println("@@@",)
	w.Write([]byte(p.body))
}

// use map instead
type multiplexer map[string]http.Handler

func (m multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := m[r.RequestURI]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

var mux = multiplexer{
	"/": home{},
}

func main() {
	http.ListenAndServe(":8080", mux)
}
*/
/*
Fake API response:
http://www.mocky.io/v2/5d99a2893100005d0097d991
http://www.mocky.io/v2/5d99a2c6310000550097d992
Fake API Failed Response (500 Error)
http://www.mocky.io/v2/5d9c2d4931000055002fc3e5
*/
