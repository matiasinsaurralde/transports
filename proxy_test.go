package transports

import (
	"testing"
)

func TestInitializationWithoutTransport(t *testing.T) {
	transport := Transport{}
	proxy := Proxy{
		Transport: transport,
	}
	proxy.Listen()
	return
}
