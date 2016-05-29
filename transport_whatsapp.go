package transports

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "errors"
)

type WhatsappTransport struct {
	*Transport
	Login         string
	Password      string
	Contact				string
	Serializer		DefaultSerializer
}

func (t *WhatsappTransport) DoLogin() bool {
	fmt.Println("FacebookTransport, Login()")
	return true

}

func (t *WhatsappTransport) Prepare() {
	fmt.Println("FacebookTransport, Prepare()")

	t.Serializer = DefaultSerializer{}
	/*
	if !t.DoLogin() {
		err := errors.New( "Authentication error!")
		panic(err)
	}
	*/

	return
}

func (t *WhatsappTransport) Handler(w http.ResponseWriter, originalRequest *http.Request) {

	client := &http.Client{}

	request, _ := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), nil)

	serializedRequest := t.Serializer.Serialize(request)

	fmt.Println("Got", originalRequest)
	fmt.Println("Serialized", string(serializedRequest))

	/*
	MessageForm, _ := t.Browser.Form("#composer_form")
	MessageForm.Input( "body", string(serializedRequest))
	MessageForm.Submit()
	*/

	resp, _ := client.Do(request)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	w.Write(b)

	return
}

func (t *WhatsappTransport) Listen() {
	fmt.Println( "FacebookTransport, Listen()")
	t.Prepare()
	fmt.Println("Polling...")
	for {
	}
	return
}
