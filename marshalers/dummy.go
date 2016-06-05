package transports

import (
	"net/http"
)

type DummyMarshaler struct {
}

func (marshaler *DummyMarshaler) Marshal(req interface{}) interface{} {
	return nil
}

func (marshaler *DummyMarshaler) DeserializeRequest(Input []byte) *http.Request {
	return nil
}
