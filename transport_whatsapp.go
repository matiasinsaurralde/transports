package transports

import (
	"fmt"
	"encoding/json"
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
	Messages			[]WhatsappMessage
}

type WhatsappMessage struct {
	Id string
	Body string
	Origin string
}

func (t *WhatsappTransport) DaemonizeWrapper() {
	fmt.Println( "WhatsappTransport, daemonizing YowsupWrapper...")
	cmd := exec.Command( "python3", YowsupHttpWrapperPath, t.Login, t.Password )
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func( t *WhatsappTransport) GetMessageIds() []string {
	MessageIds := make( []string, 0 )
	for _, Message := range t.Messages {
		MessageIds = append( MessageIds, Message.Id )
	}
	return MessageIds
}

func( t *WhatsappTransport) FetchMessages() {
	messagesUrl := strings.Join([]string{YowsupHttpWrapperUrl, "messages"}, "")
	resp, err := http.Get(messagesUrl)

	// fmt.Println( "Request:",resp, "Error:",err)

	if err != nil {
		// fmt.Println( "Wrapper error:", err)
		return
	}

	defer resp.Body.Close()

	rawBody, _ := ioutil.ReadAll( resp.Body )

	var messageList map[string]interface{}

	jsonErr := json.Unmarshal( rawBody, &messageList)

	if jsonErr != nil {
		return
	}

	MessageIds := t.GetMessageIds()

	for Id, Values := range messageList {
		ValuesMap := Values.(map[string]interface{})
		Message := WhatsappMessage{ Id: Id, Body: ValuesMap["body"].(string), Origin: ValuesMap["origin"].(string) }
		Exists := false

		for _, ExistingId := range MessageIds {
			if ExistingId == Id {
				Exists = true
				return
			}
		}

		if !Exists {
			t.Messages = append( t.Messages, Message )
		}
	}

	return
}

func (t *WhatsappTransport) DoLogin() bool {
	fmt.Println("FacebookTransport, Login()")
	return true
}

func (t *WhatsappTransport) Prepare() {
	fmt.Println("WhatsappTransport, Prepare()")

	t.Serializer = DefaultSerializer{}

	t.Messages = make([]WhatsappMessage, 0)

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
		fmt.Println( "Poll, messages:", t.Messages )
		t.FetchMessages()
		time.Sleep(1 * time.Second)
	}
	return
}
