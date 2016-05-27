package transports

import(
  "encoding/json"
  "net/http"
)

type Request struct {
	Method  string
	URL     string
	Proto   string
	Headers map[string][]string
}

type DefaultSerializer struct {
}

func( serializer *DefaultSerializer ) Serialize( req *http.Request ) []byte {
    r := Request{
      Method:  req.Method,
      URL:     req.URL.String(),
      Proto:   req.Proto,
      Headers: req.Header,
    }
    output, _ := json.Marshal(r)
    return output
}
