package transports

import (
  "log"
)

type DummyMarshaler struct {
}

func (marshaler DummyMarshaler) Marshal(i *interface{}) (error, interface{}) {
  log.Println("** DummyMarshaler, input", *i)
  var err error
	return err, []byte("aa")
}

func (marshaler DummyMarshaler) Unmarshal() {
	return
}
