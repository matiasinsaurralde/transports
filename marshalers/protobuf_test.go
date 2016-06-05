package transports

import(
	"github.com/matiasinsaurralde/transports/marshalers"
	// "github.com/matiasinsaurralde/transports/marshalers/protos"

	"net/http"
	"net/url"
	"testing"
	"log"
)

var request http.Request

func init() {

	log.Println("init")

	url, _ := url.Parse( "http://whatismyip.akamai.com/")

	request = http.Request{
		Method: "GET",
		URL: url,
		Proto: "HTTP/1.0",
	}
}

func TestMarshalRequest( t *testing.T ) {
	marshaler := transports.ProtobufMarshaler{}
	output := marshaler.Marshal( request )

	if output == nil {
		log.Fatal( "Output is nil" )
	}

	return
}
