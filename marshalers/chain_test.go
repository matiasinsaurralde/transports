package transports_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/matiasinsaurralde/transports/marshalers"
)

const ChainTestBasicChainingError string = "Couldn't match Request URL field after chaining"
const ChainTestMarshalerCountError string = "A chain should have at least two Marshalers"

func init() {

	url, _ := url.Parse("http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL:    url,
		Proto:  "HTTP/1.0",
	}
}

func TestBasicChaining(t *testing.T) {
	chain, err := transports.NewChain(
		transports.DummyMarshaler{},
		transports.DummyMarshaler{},
	)

	if err != nil {
		t.Fatal(err)
	}

	output, err := chain.Marshal(&request)

	if err != nil {
		t.Fatal(err)
	}

	outputRequest := output.(*http.Request)

	if outputRequest.URL != request.URL {
		t.Fatal(ChainTestBasicChainingError)
	}

}

func TestChainingWithSingleOrNoMarshalers(t *testing.T) {
	chain, err := transports.NewChain(transports.ProtobufMarshaler{})

	if err == nil || chain != nil {
		t.Fatal(ChainTestMarshalerCountError)
	}
	if strings.Index(err.Error(), transports.ChainSingleMarshalerError) < 0 {
		t.Fatal(ChainTestMarshalerCountError)
	}

	chain, err = transports.NewChain()
	if err == nil || chain != nil {
		t.Fatal(ChainTestMarshalerCountError)
	}
	if strings.Index(err.Error(), transports.ChainSingleMarshalerError) < 0 {
		t.Fatal(ChainTestMarshalerCountError)
	}
}
