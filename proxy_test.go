package transports

import (
	"testing"
	"fmt"
)

func TestInitializationWithoutTransport( t *testing.T ) {
	transport := Transport{}
	proxy := Proxy{
		Transport: transport,
	}
	proxy.Listen()
	return
}
