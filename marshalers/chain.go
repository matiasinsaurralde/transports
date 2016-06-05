package transports

import("log")

type ChainData struct {
  marshalers []Marshaler
  input interface{}
}

type Chain interface{
  Marshal( interface{} ) ( error, interface{} )
  process() ( error, interface{} )
}

func NewChain( marshalers ...Marshaler) Chain {
  log.Println("Got", marshalers)
  data := ChainData{marshalers, nil}
  return Chain( &data )
  /*
  c := Chain()
  for i, m := range marshalers {
    log.Println(i,m)
  }
  log.Println("x", c.Process() )
  */
}

func (s *ChainData) process() ( error, interface{} ) {
  log.Println("Process()", s )
  /*x := DummyMarshaler{}
  mm := Marshaler(&x)
  log.Println("xd",mm)*/

  var output interface{}
  // var errors bool = false

  var err error

  for i, m := range s.marshalers {
    
    log.Println( "--> Chain step #", i )
    if output == nil {
      log.Println( "No previous output, starting chain" )
      err, output = m.Marshal(&s.input)
      log.Println( "First output:", output, "Error:", err )
    } else {
      log.Println( "Previous output", output )
      err, output = m.Marshal(&output)
      log.Println( "New output:", output, "Error:", err )
    }

    if err != nil {
      break
    }

  }

  return err, output
}

func (s *ChainData) Marshal(i interface{}) ( error, interface{} ) {
  log.Println("Output()", s)
  s.input = i
  err, output := s.process()
  return err, output
}
