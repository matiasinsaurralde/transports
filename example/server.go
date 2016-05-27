package main

import(
  "github.com/matiasinsaurralde/transports"
  "github.com/joho/godotenv"
  "fmt"
  "os"
)

func main() {
  godotenv.Load()

  fmt.Println("Transports test")

  facebookTransport := transports.FacebookTransport{
    Login: os.Getenv( "FB_LOGIN" ),
    Password: os.Getenv( "FB_PASSWORD" ),
    Friend: os.Getenv( "FB_FRIEND" ),
  }

  facebookTransport.Listen()
}
