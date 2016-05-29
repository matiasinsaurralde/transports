package transports

import(
  "encoding/json"
  "net/http"
  "fmt"
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
  fmt.Println("***_", req)
    r := Request{
      Method:  req.Method,
      URL:     req.URL.String(),
      Proto:   req.Proto,
      Headers: req.Header,
    }
    fmt.Println("***r",r)
    output, _ := json.Marshal(r)
    fmt.Println("***m", string(output))
    return output
}

func( serializer *DefaultSerializer ) Deserialize( Input []byte ) *http.Request {
  r := Request{}

  json.Unmarshal( Input, &r )
  request, _ := http.NewRequest( r.Method , r.URL, nil)
  
  return request
}
