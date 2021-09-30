package fetch

type RequestBody struct {
	Urls []string
}
type Result struct {
	Responses []Response
}
type Response struct {
	Url    string
	Status int
	Body   string
}
