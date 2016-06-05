package transports

import (
	"net/http"
)

type ProtobufMarshaler struct {
}

func (marshaler *ProtobufMarshaler) Marshal(req interface{}) interface{} {
	return nil
}

func (marshaler *ProtobufMarshaler) DeserializeRequest(Input []byte) *http.Request {

	return nil
}
