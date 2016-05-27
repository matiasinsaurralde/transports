package main

import(
  "github.com/matiasinsaurralde/transports"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Transports test")

  facebookTransport := transports.FacebookTransport{
    Login: os.Getenv( "FB_LOGIN" ),
    Password: os.Getenv( "FB_PASSWORD" ),
  }

  Proxy := transports.Proxy{
    Transport: facebookTransport,
    Port: 8080,
  }

  Proxy.Listen()
}
