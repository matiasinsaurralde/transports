package main

import(
  "github.com/matiasinsaurralde/transports"
  "github.com/joho/godotenv"
  // "net/http"
  // "io/ioutil"
  "fmt"
  "os"
)

func main() {
  godotenv.Load()

  fmt.Println("Transports test (Whatsapp/Yowsup)")

  whatsappTransport := transports.WhatsappTransport{
    Login: os.Getenv( "WA_SERVER_LOGIN" ),
    Password: os.Getenv( "WA_SERVER_PASSWORD" ),
    Contact: os.Getenv( "WA_SERVER_CONTACT" ),
    YowsupWrapperPort: "8889",
  }

  whatsappTransport.Listen( func( t *transports.WhatsappTransport ) {
    // fmt.Println("callback!", t.Messages)
    for _, Value := range t.Messages {
      request := t.Serializer.Deserialize([]byte(Value.Body))
      if request.Method == "" {
        fmt.Println( "Ignoring message", Value.Id)
        t.PurgeMessage( Value.Id )
      } else {
        fmt.Println( "Accepting message", Value.Id, request)
      }
      t.PurgeMessage( Value.Id)
      /*
      client := &http.Client{}

      response, _ := client.Do(request)

      defer response.Body.Close()

      rawBody, _ := ioutil.ReadAll( response.Body)

      t.SendMessage( string(rawBody))
      t.PurgeMessage( Value.Id)
      */
    }

    t.Messages = make([]transports.WhatsappMessage, 0)
  })
}
