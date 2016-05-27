package transports

import (
	"fmt"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"errors"
)

type FacebookTransport struct {
	*Transport
	Login         string
	Password      string
	Browser       *browser.Browser
}

func (t *FacebookTransport) Prepare() {
	fmt.Println("FacebookTransport, Prepare()")
	t.Browser = surf.NewBrowser()

	if !t.DoLogin() {
		err := errors.New( "Authentication error!")
		panic(err)
	}

	return
}

func (t *FacebookTransport) DoLogin() bool {
	fmt.Println("FacebookTransport, Login()")
	err := t.Browser.Open("https://mobile.facebook.com/")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Browser.Body())
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

	return true

}
