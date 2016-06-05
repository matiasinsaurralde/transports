package transports_test

import (
	"bytes"
	"github.com/matiasinsaurralde/transports/marshalers"
	"gopkg.in/kothar/brotli-go.v0/dec"
	"strings"
	"testing"
)

func TestBrotliMarshalChain(t *testing.T) {

	_, chain := transports.NewChain(
		transports.ProtobufMarshaler{},
		transports.DummyMarshaler{},
	)
	_, output := chain.Marshal(&request)

	protobufOutput := output.([]byte)

	_, chain = transports.NewChain(
		transports.ProtobufMarshaler{},
		transports.BrotliMarshaler{},
	)

	_, output = chain.Marshal(&request)

	compressedOutput := output.([]byte)

	decompressedProtobuf, _ := dec.DecompressBuffer(compressedOutput, make([]byte, 0))

	if !bytes.Equal(protobufOutput, decompressedProtobuf) {
		t.Fatal(transports.MarshalerUnexpectedOutput)
	}

}

func TestBrotliUnsupportedType(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.BrotliMarshaler{}

	var v UnknownType
	v = UnknownType{"Value"}

	var i interface{}
	i = v

	err, _ := marshaler.Marshal(&i)

	exists := strings.Index(err.Error(), transports.MarshalerTypeNotSupportedError)

	if exists < 0 {
		t.Fatal("Unsupported type doesn't break the Brotli marshaler")
	}
}

func TestBrotliNilInput(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.BrotliMarshaler{}
	err, _ := marshaler.Marshal(nil)

	if strings.Index(err.Error(), transports.MarshalerNilTypeError) < 0 {
		t.Fatal("Nil type doesn't break the Brotli marshaler")
	}
}
