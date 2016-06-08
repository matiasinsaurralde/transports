package transports

import (
	"net/http"
	"golang.org/x/net/proxy"
)

type Transport struct {
	Name string
}

func (t *Transport) Prepare() {
	return
}

func (t *Transport) Handler(w http.ResponseWriter, req *http.Request) {
	return
}

func (t *Transport) Listen() {
	return
}

func TorDialer() (proxy.Dialer) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	if err != nil {
		panic(err)
	}
	return dialer
}
