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
	Marshal(interface{}) (interface{}, error)
	Unmarshal(interface{}) (interface{}, error)
	process(bool) (interface{}, error)
	reverse() []Marshaler
}

func NewChain(marshalers ...Marshaler) (Chain, error) {
	var err error
	if len(marshalers) <= 1 {
		err = errors.New(ChainSingleMarshalerError)
		return nil, err
	}
	data := ChainData{marshalers, nil}
	return Chain(&data), err
}

func (s *ChainData) process(reverseOrder bool) (interface{}, error) {
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
				output, err = m.Unmarshal(&s.input)
			} else {
				output, err = m.Marshal(&s.input)
			}
			log.Println("First output:", output, "Error:", err)
		} else {
			log.Println("Previous output", output)
			if reverseOrder {
				output, err = m.Unmarshal(&output)
			} else {
				output, err = m.Marshal(&output)
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

	return output, err
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

func (s *ChainData) Marshal(i interface{}) (interface{}, error) {
	log.Println("Output()", s)
	s.input = i
	output, err := s.process(false)
	return output, err
}

func (s *ChainData) Unmarshal(i interface{}) (interface{}, error) {
	log.Println("Output()", s)
	s.input = i
	output, err := s.process(true)
	return output, err
}
