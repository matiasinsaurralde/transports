package transports_test

import(
	"github.com/matiasinsaurralde/transports/marshalers"
	// "github.com/matiasinsaurralde/transports/marshalers/protos"

	"strings"
	"net/http"
	"net/url"
	"testing"
	"log"
)

type UnknownType struct {
	Field string
}

const TestRequestUrl string = "http://whatismyip.akamai.com/"

var request http.Request

func init() {

	log.Println("init")

	url, _ := url.Parse( TestRequestUrl)

	request = http.Request{
		Method: "GET",
		URL: url,
		Proto: "HTTP/1.0",
	}
}

func TestHttpRequestMarshal( t *testing.T ) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	var i interface{}
	i = request
	_, err := marshaler.Marshal(&i)
	if err != nil {
		t.Fatal("Can't marshal HttpRequest")
	}
}

func TestHttpResponseMarshal( t *testing.T ) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	var i interface{}
	i = request
	_, err := marshaler.Marshal(&i)
	if err != nil {
		t.Fatal("Can't marshal HttpResponse")
	}
}

func TestUnsupportedType( t *testing.T ) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}

	var v UnknownType
	v = UnknownType{"Value"}

	var i interface{}
	i = v

	err, _ := marshaler.Marshal(&i)

	exists := strings.Index(err.Error(), transports.MarshalerTypeNotSupportedError)

	if exists < 0 {
		t.Fatal("Unsupported type doesn't break the Protobuf marshaler")
	}
}

func TestNilInput( t *testing.T ) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	err, _ := marshaler.Marshal(nil)

	if strings.Index( err.Error(), transports.MarshalerNilTypeError ) < 0 {
		t.Fatal("Nil type doesn't break the Protobuf marshaler")
	}
}
