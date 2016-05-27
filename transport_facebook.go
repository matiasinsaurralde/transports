package transports

import (
	"fmt"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type FacebookTransport struct {
	Login         string
	Password      string
	Browser       *browser.Browser
	Authenticated bool
}

func (t *FacebookTransport) Prepare() {
	fmt.Println("FacebookTransport, Prepare()")
	t.Browser = surf.NewBrowser()
	t.Authenticated = false

	t.DoLogin()

	return
}

func (t *FacebookTransport) DoLogin() {
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

	// fmt.Println( t.Browser.Body() )
}
