package transports

import("log")

type ChainData struct {
  marshalers []interface{}
  input interface{}
}

type Chain interface{
  Process() interface{}
  Input(interface{})
  Output() ( error, interface{} )
}

func NewChain( marshalers ...interface{}) Chain {
  log.Println("Got", marshalers)
  data := ChainData{marshalers, nil}
  return Chain( data)
  /*
  c := Chain()
  for i, m := range marshalers {
    log.Println(i,m)
  }
  log.Println("x", c.Process() )
  */
}

func (s ChainData) Process() interface{} {
  log.Println("Process?", s)
  return nil
}

func (s ChainData) Input(interface{}) {
  return
}

func (s ChainData) Output() ( error, interface{} ) {
  var err error
  return err, nil
}
