package transports_test

import (
	"github.com/matiasinsaurralde/transports/marshalers"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func init() {

	url, _ := url.Parse("http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL:    url,
		Proto:  "HTTP/1.0",
	}
}

func TestBasicChaining(t *testing.T) {
	err, chain := transports.NewChain(transports.DummyMarshaler{},
		                                transports.DummyMarshaler{})

  if err != nil {
    t.Fatal(err)
  }

	err, output := chain.Marshal(&request)

  if err != nil {
    t.Fatal(err)
  }

  outputRequest := output.(*http.Request)

  if outputRequest.URL != request.URL {
    t.Fatal( "Couldn't match Request URL field after chaining")
  }

	return
}

func TestChainingWithSingleOrNoMarshalers(t *testing.T) {
	err, chain := transports.NewChain(transports.ProtobufMarshaler{})
	if err == nil || chain != nil {
		t.Fatal("A chain should have at least two Marshalers")
	}
	if strings.Index(err.Error(), transports.ChainSingleMarshalerError) < 0 {
		t.Fatal("A chain should have at least two Marshalers")
	}

	err, chain = transports.NewChain()
	if err == nil || chain != nil {
		t.Fatal("A chain should have at least two Marshalers")
	}
	if strings.Index(err.Error(), transports.ChainSingleMarshalerError) < 0 {
		t.Fatal("A chain should have at least two Marshalers")
	}
}
