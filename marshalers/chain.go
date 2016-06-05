package transports

import (
	"errors"
	"log"
)

const ChainSingleMarshalerError string = "A chain requires two or more Marshalers."
const ChainNilOutput string = "The chain returned nil"

type ChainData struct {
	marshalers []Marshaler
	input      interface{}
}

type Chain interface {
	Marshal(interface{}) (error, interface{})
	Unmarshal(interface{}) (error, interface{})
	process(bool) (error, interface{})
	reverse() []Marshaler
}

func NewChain(marshalers ...Marshaler) (error, Chain) {
	var err error
	if len(marshalers) <= 1 {
		err = errors.New(ChainSingleMarshalerError)
		return err, nil
	}
	data := ChainData{marshalers, nil}
	return err, Chain(&data)
}

func (s *ChainData) process(reverseOrder bool) (error, interface{}) {
	log.Println("Process()", s)

	var output interface{}

	var err error

	var marshalers []Marshaler

	if reverseOrder {
		marshalers = s.reverse()
	} else {
		marshalers = s.marshalers
	}

	for i, m := range marshalers {

		log.Println("--> Chain step #", i)
		if output == nil {
			log.Println("No previous output, starting chain")
      if reverseOrder {
        err, output = m.Unmarshal(&s.input)
      } else {
        err, output = m.Marshal(&s.input)
      }
			log.Println("First output:", output, "Error:", err)
		} else {
			log.Println("Previous output", output)
      if reverseOrder {
        err, output = m.Unmarshal(&output)
      } else {
        err, output = m.Marshal(&output)
      }
			log.Println("New output:", output, "Error:", err)
		}

		if err != nil {
			break
		}

	}

  if output == nil {
    err = errors.New(ChainNilOutput)
  }

	return err, output
}

func (s *ChainData) reverse() []Marshaler {
	marshalers := make([]Marshaler, len(s.marshalers))
	rindex := 0
	for i := len(s.marshalers) - 1; i >= 0; i-- {
		m := s.marshalers[i]
		marshalers[rindex] = m
		rindex++
	}
	return marshalers
}

func (s *ChainData) Marshal(i interface{}) (error, interface{}) {
	log.Println("Output()", s)
	s.input = i
	err, output := s.process(false)
	return err, output
}

func (s *ChainData) Unmarshal(i interface{}) (error, interface{}) {
	log.Println("Output()", s)
	s.input = i
	err, output := s.process(true)
	return err, output
}
