package transports

import (
// "log"
)

type DummyMarshaler struct {
}

func (marshaler DummyMarshaler) Marshal(i *interface{}) (error, interface{}) {
	var err error
	return err, *i
}

func (marshaler DummyMarshaler) Unmarshal(i *interface{}) (error, interface{}) {
	return nil, nil
}
