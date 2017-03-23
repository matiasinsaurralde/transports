package transports

import (
	"errors"

	"gopkg.in/kothar/brotli-go.v0/dec"
	"gopkg.in/kothar/brotli-go.v0/enc"
)

type BrotliMarshaler struct {
}

func (marshaler BrotliMarshaler) Marshal(i *interface{}) (interface{}, error) {
	var err error

	if i == nil {
		err = errors.New(MarshalerNilTypeError)
		return err, nil
	}

	switch (*i).(type) {
	case []byte:
	default:
		err = errors.New(MarshalerTypeNotSupportedError)
		return err, nil
	}

	var inputBuf []byte
	inputBuf = (*i).([]byte)

	buf, err := enc.CompressBuffer(nil, inputBuf, make([]byte, 0))

	return buf, err
}

func (marshaler BrotliMarshaler) Unmarshal(i *interface{}) (interface{}, error) {
	var err error

	if i == nil {
		err = errors.New(MarshalerNilTypeError)
		return err, nil
	}

	switch (*i).(type) {
	case []byte:
	default:
		err = errors.New(MarshalerTypeNotSupportedError)
		return err, nil
	}

	var inputBuf []byte
	inputBuf = (*i).([]byte)

	buf, err := dec.DecompressBuffer(inputBuf, make([]byte, 0))

	return buf, err
}
