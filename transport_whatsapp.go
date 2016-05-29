package transports

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os/exec"
	"time"
	"strings"
	// "errors"
)

const YowsupHttpWrapperPath = "../yowsup-http-wrapper/run.py"
const YowsupHttpWrapperUrl = "http://127.0.0.1:8888/"

type WhatsappTransport struct {
	*Transport
	Login         string
	Password      string
	Contact				string
	Serializer		DefaultSerializer
}

func (t *WhatsappTransport) DaemonizeWrapper() {
	fmt.Println( "WhatsappTransport, daemonizing YowsupWrapper...")
	cmd := exec.Command( "python3", YowsupHttpWrapperPath, t.Login, t.Password )
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func( t *WhatsappTransport) FetchMessages() {
	// messagesUrl := strings.Join(%{YowsupHttpWrapperUrl})
	// resp, err := http.Get()
	/fmt.Println(123,messagesUrl)
	return
}

func (t *WhatsappTransport) DoLogin() bool {
	fmt.Println("FacebookTransport, Login()")
	return true
}

func (t *WhatsappTransport) Prepare() {
	fmt.Println("WhatsappTransport, Prepare()")

	t.Serializer = DefaultSerializer{}

	go t.DaemonizeWrapper()

	go t.Listen()

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
	fmt.Println("Polling...")
	for {
		fmt.Println( "Poll." )
		t.FetchMessages()
		time.Sleep(5 * time.Second)
	}
	return
}
