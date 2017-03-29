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

const YowsupHTTPWrapperPath = "../yowsup-http-wrapper/run.py"

var ResponseChannel chan Response

type WhatsappTransport struct {
	*Transport
	Login             string
	Password          string
	Contact           string
	YowsupWrapperPort string
	YowsupWrapperURL  string
	UseTor            bool
	Serializer        DefaultSerializer
	Messages          []WhatsappMessage
	// ResponseChannel	chan Response
}

type WhatsappMessage struct {
	ID     string `json:"id,omitempty"`
	Body   string `json:"msg,omitempty"`
	Origin string `json:"origin,omitempty"`
	Dest   string `json:"dest,omitempty"`
}

type WhatsappMessageCallback func(*WhatsappTransport)

func (t *WhatsappTransport) DaemonizeWrapper() {
	fmt.Println("WhatsappTransport, daemonizing YowsupWrapper...")

	t.YowsupWrapperURL = fmt.Sprintf("http://127.0.0.1:%s/", t.YowsupWrapperPort)

	cmd := exec.Command("python3", YowsupHTTPWrapperPath, t.Login, t.Password, t.YowsupWrapperPort)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func (t *WhatsappTransport) GetMessageIDs() []string {
	ids := make([]string, 0)
	for _, message := range t.Messages {
		ids = append(ids, message.ID)
	}
	return ids
}

func (t *WhatsappTransport) PurgeMessage(Id string) {
	messagesURL := fmt.Sprintf("%s%s?id=%s", t.YowsupWrapperURL, "messages", Id)
	deleteRequest, _ := http.NewRequest("DELETE", messagesURL, nil)
	http.DefaultClient.Do(deleteRequest)
}

func (t *WhatsappTransport) FetchMessages() {
	messagesURL := strings.Join([]string{t.YowsupWrapperURL, "messages"}, "")
	resp, err := http.Get(messagesURL)

	if err != nil {
		// fmt.Println( "Wrapper error:", err)
		return
	}

	defer resp.Body.Close()

	rawBody, _ := ioutil.ReadAll(resp.Body)

	var messageList map[string]interface{}

	err = json.Unmarshal(rawBody, &messageList)

	if err != nil {
		return
	}

	messageIDs := t.GetMessageIDs()

	for id, values := range messageList {
		valuesMap := values.(map[string]interface{})
		message := WhatsappMessage{ID: id, Body: valuesMap["body"].(string), Origin: valuesMap["origin"].(string)}
		exists := false

		for _, existingID := range messageIDs {
			if existingID == id {
				exists = true
				return
			}
		}

		if !exists {
			t.Messages = append(t.Messages, message)
		}
	}
}

func (t *WhatsappTransport) SendMessage(body string) {
	messagesURL := strings.Join([]string{t.YowsupWrapperURL, "messages"}, "")
	message := WhatsappMessage{Body: body, Dest: t.Contact}
	jsonBuffer, _ := json.Marshal(&message)
	http.Post(messagesURL, "application/json", bytes.NewReader(jsonBuffer))

	fmt.Println("<-- Sending message\n", message)

	return
}

func (t *WhatsappTransport) Prepare() {
	// fmt.Println("WhatsappTransport, Prepare()")

	t.YowsupWrapperURL = fmt.Sprintf("http://127.0.0.1:%s/", t.YowsupWrapperPort)

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

	fmt.Println("--> Receiving message\n", response)
	fmt.Println("<-> Forwarding message\n", response)

	w.Write([]byte(response.Body))

}

func (t *WhatsappTransport) HandleClientMessages() {
	for _, Value := range t.Messages {

		response := t.Serializer.DeserializeResponse([]byte(Value.Body))

		t.PurgeMessage(Value.ID)

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
}
