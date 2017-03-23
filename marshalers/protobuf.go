package transports

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/matiasinsaurralde/transports/marshalers/protos"
)

type ProtobufMarshaler struct {
}

func (marshaler ProtobufMarshaler) Marshal(i *interface{}) (interface{}, error) {
	var err error
	var r interface{}

	if i == nil {
		err = errors.New(MarshalerNilTypeError)
		return r, err
	}

	switch t := (*i).(type) {
	case *http.Request:
		request := (*i).(*http.Request)
		requestProto := &transportsProto.HttpRequest{
			Method: proto.String(request.Method),
			Url:    proto.String(request.URL.String()),
			Proto:  proto.String(request.Proto),
		}
		r, err = proto.Marshal(requestProto)
	case *http.Response:
	default:
		message := fmt.Sprintf(MarshalerTypeNotSupportedError)
		typestr := fmt.Sprintf("%T", t)
		err = errors.New(strings.Join([]string{message, typestr}, " "))
	}

	return r, err
}

func (marshaler ProtobufMarshaler) Unmarshal(i *interface{}) (interface{}, error) {
	var err error
	var r interface{}

	switch t := (*i).(type) {
	case []byte:
	default:
		message := fmt.Sprintf(MarshalerTypeNotSupportedError)
		typestr := fmt.Sprintf("%T", t)
		err = errors.New(strings.Join([]string{message, typestr}, " "))
	}

	buffer := (*i).([]byte)

	httpResponse := &transportsProto.HttpResponse{}
	err = proto.Unmarshal(buffer, httpResponse)

	if err == nil {
		r = httpResponse
		return r, err
	}

	httpRequest := &transportsProto.HttpRequest{}
	err = proto.Unmarshal(buffer, httpRequest)

	if err == nil {
		r = httpRequest
		return r, err
	}

	return err, nil
}
