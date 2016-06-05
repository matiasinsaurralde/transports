package transports_test

import (
	"github.com/matiasinsaurralde/transports/marshalers"
	// "github.com/matiasinsaurralde/transports/marshalers/protos"

	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func init() {

	log.Println("init")

	url, _ := url.Parse("http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL:    url,
		Proto:  "HTTP/1.0",
	}
}

func TestBasicChaining(t *testing.T) {
	err, chain := transports.NewChain(transports.ProtobufMarshaler{},
		transports.DummyMarshaler{})

	// chain := transports.Chain()
	// chain.process()
	err, output := chain.Marshal(&request)
	log.Println("err", err, "output", output)
	// log.Println(1,chain)

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
