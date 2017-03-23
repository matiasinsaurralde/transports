package transports

const MarshalerNilTypeError string = "Marshaler can't handle a nil value."
const MarshalerTypeNotSupportedError string = "Marshaler doesn't support the type you're using:"

const MarshalerUnexpectedOutput string = "Unexpected Marshaler output"

type Marshaler interface {
	Marshal(*interface{}) (interface{}, error)
	Unmarshal(*interface{}) (interface{}, error)
}
