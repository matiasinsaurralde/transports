package transports

import (
	"net/http"
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
