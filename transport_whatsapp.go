package transports

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
	// "errors"
)

const YowsupHttpWrapperPath = "../yowsup-http-wrapper/run.py"

var ResponseChannel chan Response

type WhatsappTransport struct {
	*Transport
	Login             string
	Password          string
	Contact           string
	YowsupWrapperPort string
	YowsupWrapperUrl  string
	Serializer        DefaultSerializer
	Messages          []WhatsappMessage
	// ResponseChannel	chan Response
}

type WhatsappMessage struct {
	Id     string `json:"id,omitempty"`
	Body   string `json:"msg,omitempty"`
	Origin string `json:"origin,omitempty"`
	Dest   string `json:"dest,omitempty"`
}

type WhatsappMessageCallback func(*WhatsappTransport)

func (t *WhatsappTransport) DaemonizeWrapper() {
	fmt.Println("WhatsappTransport, daemonizing YowsupWrapper...")

	t.YowsupWrapperUrl = fmt.Sprintf("http://127.0.0.1:%s/", t.YowsupWrapperPort)

	cmd := exec.Command("python3", YowsupHttpWrapperPath, t.Login, t.Password, t.YowsupWrapperPort)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func (t *WhatsappTransport) GetMessageIds() []string {
	MessageIds := make([]string, 0)
	for _, Message := range t.Messages {
		MessageIds = append(MessageIds, Message.Id)
	}
	return MessageIds
}

func (t *WhatsappTransport) PurgeMessage(Id string) {
	messagesUrl := fmt.Sprintf("%s%s?id=%s", t.YowsupWrapperUrl, "messages", Id)
	deleteRequest, _ := http.NewRequest("DELETE", messagesUrl, nil)
	http.DefaultClient.Do(deleteRequest)
}

func (t *WhatsappTransport) FetchMessages() {
	messagesUrl := strings.Join([]string{t.YowsupWrapperUrl, "messages"}, "")
	resp, err := http.Get(messagesUrl)

	if err != nil {
		// fmt.Println( "Wrapper error:", err)
		return
	}

	defer resp.Body.Close()

	rawBody, _ := ioutil.ReadAll(resp.Body)

	var messageList map[string]interface{}

	jsonErr := json.Unmarshal(rawBody, &messageList)

	if jsonErr != nil {
		return
	}

	MessageIds := t.GetMessageIds()

	for Id, Values := range messageList {
		ValuesMap := Values.(map[string]interface{})
		Message := WhatsappMessage{Id: Id, Body: ValuesMap["body"].(string), Origin: ValuesMap["origin"].(string)}
		Exists := false

		for _, ExistingId := range MessageIds {
			if ExistingId == Id {
				Exists = true
				return
			}
		}

		if !Exists {
			t.Messages = append(t.Messages, Message)
		}
	}

	return
}

func (t *WhatsappTransport) SendMessage(body string) {
	messagesUrl := strings.Join([]string{t.YowsupWrapperUrl, "messages"}, "")
	message := WhatsappMessage{Body: body, Dest: t.Contact}
	jsonBuffer, _ := json.Marshal(&message)
	http.Post(messagesUrl, "application/json", bytes.NewReader(jsonBuffer))

	fmt.Println("<-- Sending message\n", message, "\n")

	return
}

func (t *WhatsappTransport) Prepare() {
	// fmt.Println("WhatsappTransport, Prepare()")

	t.YowsupWrapperUrl = fmt.Sprintf("http://127.0.0.1:%s/", t.YowsupWrapperPort)

	t.Serializer = DefaultSerializer{}

	t.Messages = make([]WhatsappMessage, 0)

	ResponseChannel = make(chan Response)

	return
}

func (t *WhatsappTransport) Handler(w http.ResponseWriter, originalRequest *http.Request) {

	// client := &http.Client{}

	// request, _ := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), nil)

	serializedRequest := t.Serializer.Serialize(originalRequest, true).([]byte)

	go t.SendMessage(string(serializedRequest))

	w.Header().Set("Via", fmt.Sprintf("WhatsappTransport/%s", t.Contact))

	response := <-ResponseChannel

	for HeaderKey, HeaderValue := range response.Headers {
		w.Header().Set(HeaderKey, HeaderValue[0])
	}

	fmt.Println("--> Receiving message\n", response, "\n")
	fmt.Println("<-> Forwarding message\n", response, "\n")

	w.Write([]byte(response.Body))

}

func (t *WhatsappTransport) HandleClientMessages() {
	for _, Value := range t.Messages {

		response := t.Serializer.DeserializeResponse([]byte(Value.Body))

		t.PurgeMessage(Value.Id)

		go func(r Response) {
			ResponseChannel <- response
		}(response)

	}
	t.Messages = make([]WhatsappMessage, 0)
}

func (t *WhatsappTransport) Listen(Callback WhatsappMessageCallback) {

	// fmt.Println( "WhatsappTransport, Listen()")
	fmt.Println("Polling...")

	if Callback != nil {
		t.Prepare()
	}

	go t.DaemonizeWrapper()

	for {
		// fmt.Println( "Poll, messages:", t.Messages )
		t.FetchMessages()

		if Callback == nil {
			t.HandleClientMessages()
		} else {
			Callback(t)
		}
		time.Sleep(1 * time.Second)
	}
	return
}
