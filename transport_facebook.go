package transports

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"errors"
)

type FacebookTransport struct {
	*Transport
	Login         string
	Password      string
	Friend				string
	Browser       *browser.Browser
	Serializer		DefaultSerializer
	ChatUrl string
}

func (t *FacebookTransport) DoLogin() bool {
	fmt.Println("FacebookTransport, Login()")
	err := t.Browser.Open("https://mobile.facebook.com/")
	if err != nil {
		panic(err)
	}

	LoginForm := t.Browser.Forms()[1]
	LoginForm.Input("email", t.Login)
	LoginForm.Input("pass", t.Password)
	if LoginForm.Submit() != nil {
		panic(err)
	}

	err = t.Browser.Open("https://mobile.facebook.com/profile.php")

	if err != nil {
		panic(err)
	}

	fmt.Println("Logged in as", t.Browser.Title(), "?")

	FriendUrl := fmt.Sprintf("https://mobile.facebook.com/%s", t.Friend )
	err = t.Browser.Open( FriendUrl)

	if err != nil {
		panic(err)
	}

	t.Browser.Click("a[href*=\"/messages/thread/\"]")

	t.ChatUrl = t.Browser.Url().String()

	return true

}

func (t *FacebookTransport) Prepare() {
	fmt.Println("FacebookTransport, Prepare()")

	t.Serializer = DefaultSerializer{}

	t.Browser = surf.NewBrowser()

	if !t.DoLogin() {
		err := errors.New( "Authentication error!")
		panic(err)
	}

	return
}

func (t *FacebookTransport) Handler(w http.ResponseWriter, originalRequest *http.Request) {

	t.Browser.Open( t.ChatUrl )

	client := &http.Client{}

	request, _ := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), nil)

	serializedRequest := t.Serializer.Serialize(request)

	fmt.Println("Got", originalRequest)
	fmt.Println("Serialized", string(serializedRequest))

	MessageForm, _ := t.Browser.Form("#composer_form")
	MessageForm.Input( "body", string(serializedRequest))
	MessageForm.Submit()

	resp, _ := client.Do(request)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	w.Write(b)

	return
}

func (t *FacebookTransport) Listen() {
	fmt.Println( "FacebookTransport, Listen()")
	t.Prepare()
	fmt.Println("Polling...")
	for {
	}
	return
}
