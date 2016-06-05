package transports

type Marshaler interface{
  Marshal(*interface{}) (error, interface{})
  Unmarshal()
}
