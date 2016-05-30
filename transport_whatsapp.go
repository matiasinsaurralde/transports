package transports

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"os/exec"
	"time"
	"strings"
	"bytes"
	// "errors"
)

const YowsupHttpWrapperPath = "../yowsup-http-wrapper/run.py"

type WhatsappTransport struct {
	*Transport
	Login         string
	Password      string
	Contact				string
	YowsupWrapperPort	string
	YowsupWrapperUrl string
	Serializer		DefaultSerializer
	Messages			[]WhatsappMessage
}

type WhatsappMessage struct {
	Id string	`json:"id,omitempty"`
	Body string `json:"msg,omitempty"`
	Origin string	`json:"origin,omitempty"`
	Dest string `json:"dest,omitempty"`
}

type WhatsappMessageCallback func(*WhatsappTransport)

func (t *WhatsappTransport) DaemonizeWrapper() {
	fmt.Println( "WhatsappTransport, daemonizing YowsupWrapper...")

	t.YowsupWrapperUrl = fmt.Sprintf("http://127.0.0.1:%s/", t.YowsupWrapperPort)

	cmd := exec.Command( "python3", YowsupHttpWrapperPath, t.Login, t.Password, t.YowsupWrapperPort )
	err := cmd.Run()

	fmt.Println(cmd,err)

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

func( t *WhatsappTransport) PurgeMessage( Id string ) {
	messagesUrl := fmt.Sprintf("%s%s?id=%s", t.YowsupWrapperUrl, "messages", Id)
	deleteRequest, _ := http.NewRequest( "DELETE", messagesUrl, nil)
	http.DefaultClient.Do(deleteRequest)
}

func( t *WhatsappTransport) FetchMessages() {
	messagesUrl := strings.Join([]string{t.YowsupWrapperUrl, "messages"}, "")
	resp, err := http.Get(messagesUrl)

	// fmt.Println( "Request:",resp, "Error:",err)

	if err != nil {
		fmt.Println( "Wrapper error:", err)
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

func (t *WhatsappTransport) SendMessage(body string) {
	messagesUrl := strings.Join([]string{t.YowsupWrapperUrl, "messages"}, "")
	message := WhatsappMessage{Body: body, Dest: t.Contact}
	fmt.Println("Sending message", message)
	jsonBuffer, _ := json.Marshal(&message)
	http.Post(messagesUrl, "application/json", bytes.NewReader(jsonBuffer) )
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

	// go t.DaemonizeWrapper()

	// go t.Listen(nil)

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

	serializedRequest := t.Serializer.Serialize(originalRequest)

	t.SendMessage(string(serializedRequest))

	resp, _ := client.Do(request)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	w.Write(b)

	return
}

func (t *WhatsappTransport) Listen( Callback WhatsappMessageCallback ) {

	fmt.Println( "FacebookTransport, Listen()")
	fmt.Println("Polling...")

	t.Prepare()

	go t.DaemonizeWrapper()

	for {
		fmt.Println( "Poll, messages:", t.Messages )
		t.FetchMessages()
		if Callback == nil {
		} else {
			Callback( t )
		}
		t.FetchMessages()
		time.Sleep(1 * time.Second)
	}
	return
}
