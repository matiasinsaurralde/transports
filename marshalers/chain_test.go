package transports_test

import(
	"github.com/matiasinsaurralde/transports/marshalers"
	// "github.com/matiasinsaurralde/transports/marshalers/protos"

	"net/http"
	"net/url"
	"testing"
	"log"
)

func init() {

	log.Println("init")

	url, _ := url.Parse( "http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL: url,
		Proto: "HTTP/1.0",
	}
}

func TestChaining( t *testing.T ) {
	marshaler := transports.ProtobufMarshaler{}
	output := marshaler.Marshal( request )

  if output == nil {
  }

  chain := transports.NewChain( transports.ProtobufMarshaler{},
                                transports.DummyMarshaler{} )

  chain.Input(&request)
  // chain := transports.Chain()
  output, err := chain.Output()
  log.Println("output", output, "err", err)
  // log.Println(1,chain)

	return
}
