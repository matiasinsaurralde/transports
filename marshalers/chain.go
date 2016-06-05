package transports

import (
	"errors"
	"log"
)

const ChainSingleMarshalerError string = "A chain requires two or more Marshalers."

type ChainData struct {
	marshalers []Marshaler
	input      interface{}
}

type Chain interface {
	Marshal(interface{}) (error, interface{})
	process() (error, interface{})
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

func (s *ChainData) process() (error, interface{}) {
	log.Println("Process()", s)

	var output interface{}

	var err error

	for i, m := range s.marshalers {

		log.Println("--> Chain step #", i)
		if output == nil {
			log.Println("No previous output, starting chain")
			err, output = m.Marshal(&s.input)
			log.Println("First output:", output, "Error:", err)
		} else {
			log.Println("Previous output", output)
			err, output = m.Marshal(&output)
			log.Println("New output:", output, "Error:", err)
		}

		if err != nil {
			break
		}

	}

	return err, output
}

func (s *ChainData) Marshal(i interface{}) (error, interface{}) {
	log.Println("Output()", s)
	s.input = i
	err, output := s.process()
	return err, output
}
