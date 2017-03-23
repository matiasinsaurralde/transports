package transports

type DummyMarshaler struct {
}

func (marshaler DummyMarshaler) Marshal(i *interface{}) (interface{}, error) {
	return *i, nil
}

func (marshaler DummyMarshaler) Unmarshal(i *interface{}) (interface{}, error) {
	return nil, nil
}
