package transports_test

import (
	"testing"

	"github.com/matiasinsaurralde/transports/marshalers"
)

type UnknownType struct {
	Field string
}

var TestVariable UnknownType

const TestVariableString string = "Hello world"

var TestInterface interface{}

func init() {
	TestVariable = UnknownType{TestVariableString}
	TestInterface = TestVariable
}

func TestMarshal(t *testing.T) {
	var marshaler transports.Marshaler

	marshaler = transports.DummyMarshaler{}

	output, err := marshaler.Marshal(&TestInterface)

	if err != nil {
		t.Fatal(err)
	}

	outputVariable := output.(UnknownType)

	if outputVariable.Field != TestVariableString {
		t.Fatal("Couldn't match the UnknownType field value")
	}
}
