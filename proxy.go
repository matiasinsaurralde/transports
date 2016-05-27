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

type Proxy struct {
	Port      int
	Transport interface{}
}

func (proxy *Proxy) Listen() {
	fmt.Println("Listening on", proxy.Port, ", transport:", proxy.Transport)
	transport := proxy.Transport.(FacebookTransport)
	transport.Prepare()
	return
}

func MarshalRequest(request *http.Request) []byte {
	r := Request{
		Method:  request.Method,
		URL:     request.URL.String(),
		Proto:   request.Proto,
		Headers: request.Header,
	}
	output, _ := json.Marshal(r)
	return output
}

func HandleRequest(w http.ResponseWriter, originalRequest *http.Request) {
	client := &http.Client{}
	request, _ := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), nil)

	fmt.Println("Got", originalRequest)
	fmt.Println("Recreated", request)
	fmt.Println("client", client)

	// fmt.Println(resp.Body, err)
	return
}

/*
func main() {
  fmt.Println( "Transports Test")
  // transport := &http.Transport{}
  // client := &http.Client{ Transport: transport }

  request, _ := http.NewRequest( "GET", "http://whatismyip.akamai.com/", nil)
  fmt.Println(request)

  jsonRequest := MarshalRequest(request)
  fmt.Println( "JSON Request:", string(jsonRequest))

  http.HandleFunc("/", HandleRequest)
  http.ListenAndServe(":8080", nil)
}
*/
