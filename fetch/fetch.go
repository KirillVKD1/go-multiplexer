package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var maxUrlsAmount, timeoutSeconds = 20, 1

func FetchAll(w http.ResponseWriter, req *http.Request) {
	var parsedRequestBody *RequestBody
	err := json.NewDecoder(req.Body).Decode(&parsedRequestBody)
	if err != nil {
		http.Error(w, "Body parsing error", http.StatusBadRequest)
	}

	urlsListLength := len(parsedRequestBody.Urls)

	if urlsListLength > maxUrlsAmount {
		http.Error(
			w,
			strings.Join([]string{"Max commands amount should be no more than", strconv.Itoa(maxUrlsAmount)}, " "),
			http.StatusBadRequest,
		)
		return
	}

	results := make(chan Response)
	fail := make(chan Response)

	for i := 0; i < urlsListLength; i++ {
		go Send(parsedRequestBody.Urls[i], results, fail)
	}

	// Collect results:
	responses := make([]Response, urlsListLength)

	for i := 0; i < urlsListLength; {
		select {
		case result := <-results:
			responses[i] = result
			i++
		case f := <-fail:
			i = urlsListLength
			http.Error(
				w,
				strings.Join([]string{"Url", f.Url, "failed with status:", strconv.Itoa(f.Status), "Body", f.Body}, " "),
				http.StatusBadRequest,
			)
			return
		default:
		}
	}

	jsonResponse, err := json.Marshal(Result{Responses: responses})
	if err != nil {
		http.Error(w, "Response parsing error", http.StatusBadRequest)
	}

	w.Write(jsonResponse)
}

func Send(url string, results chan<- Response, fail chan<- Response) {

	client := http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}

	res, err := client.Get(url)

	if err != nil {
		fail <- Response{
			Url:    url,
			Status: http.StatusInternalServerError,
		}
		return
	}

	status := res.StatusCode
	var strBody string
	body, err := io.ReadAll(res.Body)

	if err != nil {
		strBody = "{}"
	} else {
		strBody = string(body)
	}

	if status >= 300 {
		fail <- Response{
			Url:    url,
			Status: status,
			Body:   strBody,
		}
		return
	}

	results <- Response{Url: url, Status: status, Body: strBody}
}
