package transports_test

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/matiasinsaurralde/transports"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

var whatsappTransportClient transports.WhatsappTransport
var whatsappTransportServer transports.WhatsappTransport
var Proxy transports.Proxy

const ExternalTestUrl string = "http://whatismyip.akamai.com/"
const ProxyUrl string = "http://127.0.0.1:8080"
const ProxyPort int = 8080

func startClient() {

	// godotenv.Load()

	fmt.Println("Transports test (Whatsapp/Yowsup)")

	whatsappTransportClient = transports.WhatsappTransport{
		Login:             os.Getenv("WA_CLIENT_LOGIN"),
		Password:          os.Getenv("WA_CLIENT_PASSWORD"),
		Contact:           os.Getenv("WA_CLIENT_CONTACT"),
		YowsupWrapperPort: "8888",
	}

	go whatsappTransportClient.Listen(nil)

	Proxy = transports.Proxy{
		Transport: whatsappTransportClient,
		Port:      ProxyPort,
	}

	Proxy.Listen()
}

func startServer() {

	// godotenv.Load()

	whatsappTransportServer = transports.WhatsappTransport{
		Login:             os.Getenv("WA_SERVER_LOGIN"),
		Password:          os.Getenv("WA_SERVER_PASSWORD"),
		Contact:           os.Getenv("WA_SERVER_CONTACT"),
		YowsupWrapperPort: "8889",
	}

	whatsappTransportServer.Listen(func(t *transports.WhatsappTransport) {
		fmt.Println("callback", t)
		for _, Value := range t.Messages {
			request := t.Serializer.DeserializeRequest([]byte(Value.Body))
			if request.Method == "" {
				fmt.Println("*** Ignoring message", "\n")
				t.PurgeMessage(Value.Id)
				return
			}

			fmt.Println("--> Receiving, accepting message\n", request, "\n")
			client := &http.Client{}
			response, _ := client.Do(request)
			defer response.Body.Close()

			rawBody, _ := ioutil.ReadAll(response.Body)

			serializedResponse := t.Serializer.Serialize(response, false).(transports.Response)
			serializedResponse.Body = string(rawBody)

			jsonResponse, _ := json.Marshal(serializedResponse)

			t.SendMessage(string(jsonResponse))

			t.PurgeMessage(Value.Id)

		}

		t.Messages = make([]transports.WhatsappMessage, 0)
	})
}

func init() {
	godotenv.Load()
	go startServer()
	go startClient()
}

func testRequest() []byte {
	resp, err := http.Get(ExternalTestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func transportRequest() ([]byte, error) {
	url, _ := url.Parse(ProxyUrl)
	tr := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(ExternalTestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func TestMessageExchange(t *testing.T) {

	testBody := testRequest()

	time.Sleep(10 * time.Second)

	testTransportBody, err := transportRequest()

	if err == nil {

	} else {
		fmt.Println("testBody", testBody)
		fmt.Println("testTransportBody", testTransportBody)
	}
	for {

	}
}
