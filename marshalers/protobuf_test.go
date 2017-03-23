package transports_test

import (
	"github.com/matiasinsaurralde/transports/marshalers"
	// "github.com/matiasinsaurralde/transports/marshalers/protos"

	"net/http"
	"net/url"
	"strings"
	"testing"
)

var request http.Request

func init() {

	url, _ := url.Parse("http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL:    url,
		Proto:  "HTTP/1.0",
	}
}

func TestProtobufHttpRequestMarshal(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	var i interface{}
	i = request
	_, err := marshaler.Marshal(&i)
	if err != nil {
		t.Fatal("Can't marshal HttpRequest")
	}
}

func TestProtobufHttpResponseMarshal(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	var i interface{}
	i = request
	_, err := marshaler.Marshal(&i)
	if err != nil {
		t.Fatal("Can't marshal HttpResponse")
	}
}

func TestProtobufUnsupportedType(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}

	var v UnknownType
	v = UnknownType{"Value"}

	var i interface{}
	i = v

	_, err := marshaler.Marshal(&i)

	exists := strings.Index(err.Error(), transports.MarshalerTypeNotSupportedError)

	if exists < 0 {
		t.Fatal("Unsupported type doesn't break the Protobuf marshaler")
	}
}

func TestProtobufNilInput(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.ProtobufMarshaler{}
	_, err := marshaler.Marshal(nil)

	if strings.Index(err.Error(), transports.MarshalerNilTypeError) < 0 {
		t.Fatal("Nil type doesn't break the Protobuf marshaler")
	}
}
