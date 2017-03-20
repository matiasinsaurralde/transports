package transports

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Method  string
	URL     string
	Proto   string
	Headers map[string][]string
}

type Response struct {
	Status     string
	StatusCode int
	Proto      string
	Headers    map[string][]string
	Body       string
}

type DefaultSerializer struct {
}

func (serializer *DefaultSerializer) Serialize(req interface{}, jsonOutput bool) interface{} {

	var output []byte
	var r interface{}

	switch t := req.(type) {
	case *http.Request:
		req := req.(*http.Request)
		r = Request{
			Method:  req.Method,
			URL:     req.URL.String(),
			Proto:   req.Proto,
			Headers: req.Header,
		}
	case *http.Response:
		res := req.(*http.Response)
		r = Response{
			Status:     res.Status,
			StatusCode: res.StatusCode,
			Proto:      res.Proto,
			Headers:    res.Header,
		}
	default:
		fmt.Println("Unknown Type", t)
	}

	if jsonOutput {
		output, _ = json.Marshal(r)
		return output
	}
	return r
}

func (serializer *DefaultSerializer) DeserializeRequest(Input []byte) *http.Request {
	r := Request{}

	json.Unmarshal(Input, &r)
	request, _ := http.NewRequest(r.Method, r.URL, nil)

	return request
}

func (serializer *DefaultSerializer) DeserializeResponse(Input []byte) Response {
	r := Response{}

	json.Unmarshal(Input, &r)

	return r
}
