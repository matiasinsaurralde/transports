package transports

import(
  "github.com/headzoo/surf"
  "github.com/headzoo/surf/browser"
  "fmt"
)

type FacebookTransport struct {
  Login string
  Password string
  Browser *browser.Browser
  Authenticated bool
}

func( t *FacebookTransport ) Prepare() {
  fmt.Println("FacebookTransport, Prepare()")
  t.Browser = surf.NewBrowser()
  t.Authenticated = false

  t.DoLogin()

  return
}

func( t *FacebookTransport ) DoLogin() {
  fmt.Println("FacebookTransport, Login()")
  err := t.Browser.Open( "https://mobile.facebook.com/")
  if err != nil {
    panic(err)
  }
  LoginForm, _ := t.Browser.Form( "mobile-login-form" )
  fmt.Println(LoginForm)
  LoginForm.Input( "username", "abc")
  LoginForm.Input( "password", "123")
  LoginForm.Submit()
}
