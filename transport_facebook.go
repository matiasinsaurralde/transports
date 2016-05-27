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
}

func( t *FacebookTransport ) Prepare() {
  t.Browser = surf.NewBrowser()
  fmt.Println("FacebookTransport, Prepare()")
  return
}
