package transports_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matiasinsaurralde/transports/marshalers"
	"github.com/matiasinsaurralde/transports/marshalers/protos"
	"gopkg.in/kothar/brotli-go.v0/dec"
)

func TestBrotliMarshalChain(t *testing.T) {

	chain, _ := transports.NewChain(
		transports.ProtobufMarshaler{},
		transports.DummyMarshaler{},
	)
	output, _ := chain.Marshal(&request)

	protobufOutput := output.([]byte)

	compressionChain, _ := transports.NewChain(
		transports.ProtobufMarshaler{},
		transports.BrotliMarshaler{},
	)

	output, _ = compressionChain.Marshal(&request)

	compressedOutput := output.([]byte)

	decompressedProtobuf, _ := dec.DecompressBuffer(compressedOutput, make([]byte, 0))

	if !bytes.Equal(protobufOutput, decompressedProtobuf) {
		t.Fatal(transports.MarshalerUnexpectedOutput)
	}

}

func TestBrotliUnmarshalChain(t *testing.T) {

	chain, _ := transports.NewChain(
		transports.ProtobufMarshaler{},
		transports.BrotliMarshaler{},
	)

	output, _ := chain.Marshal(&request)

	compressedOutput := output.([]byte)

	var i interface{}
	i = compressedOutput

	unmarshalOutput, err := chain.Unmarshal(i)

	if err != nil {
		t.Fatal(err)
	}

	var protoOutput *transportsProto.HttpRequest
	protoOutput = unmarshalOutput.(*transportsProto.HttpRequest)

	if protoOutput.GetUrl() != "http://whatismyip.akamai.com/" {
		t.Fatal("Protobuffer field doesn't have the original HttpRequest value")
	}

}

func TestBrotliUnsupportedType(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.BrotliMarshaler{}

	var v UnknownType
	v = UnknownType{"Value"}

	var i interface{}
	i = v

	_, err := marshaler.Marshal(&i)

	exists := strings.Index(err.Error(), transports.MarshalerTypeNotSupportedError)

	if exists < 0 {
		t.Fatal("Unsupported type doesn't break the Brotli marshaler")
	}
}

func TestBrotliNilInput(t *testing.T) {
	var marshaler transports.Marshaler
	marshaler = transports.BrotliMarshaler{}
	_, err := marshaler.Marshal(nil)

	if strings.Index(err.Error(), transports.MarshalerNilTypeError) < 0 {
		t.Fatal("Nil type doesn't break the Brotli marshaler")
	}
}
