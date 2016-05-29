package main

import(
  "github.com/matiasinsaurralde/transports"
  "github.com/joho/godotenv"
  "fmt"
  "os"
)

func main() {
  godotenv.Load()

  fmt.Println("Transports test (Whatsapp/Yowsup)")

  whatsappTransport := transports.WhatsappTransport{
    Login: os.Getenv( "WA_LOGIN" ),
    Password: os.Getenv( "WA_PASSWORD" ),
    Contact: os.Getenv( "WA_CONTACT" ),
  }

  Proxy := transports.Proxy{
    Transport: whatsappTransport,
    Port: 8080,
  }

  Proxy.Listen()
}
