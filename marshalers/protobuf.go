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

func (marshaler ProtobufMarshaler) Marshal(i *interface{}) (error, interface{}) {
	var err error
	var r interface{}

	if i == nil {
		err = errors.New(MarshalerNilTypeError)
		return err, r
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

	return err, r
}

func (marshaler ProtobufMarshaler) Unmarshal(i *interface{}) (error, interface{}) {
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

	HttpResponse := &transportsProto.HttpResponse{}
	err = proto.Unmarshal(buffer, HttpResponse)

	if err == nil {
		r = HttpResponse
		return err, r
	}

	HttpRequest := &transportsProto.HttpRequest{}
	err = proto.Unmarshal(buffer, HttpRequest)

	if err == nil {
		r = HttpRequest
		return err, r
	}

	return err, nil
}
