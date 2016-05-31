package main

import(
  "github.com/matiasinsaurralde/transports"
  "github.com/joho/godotenv"
  "encoding/json"
  "net/http"
  "io/ioutil"
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
      request := t.Serializer.DeserializeRequest([]byte(Value.Body))
      if request.Method == "" {
        fmt.Println( "Ignoring message", Value.Id)
        t.PurgeMessage( Value.Id )
        return
      }

      fmt.Println( "Accepting message", Value.Id, request)
      client := &http.Client{}
      response, _ := client.Do( request)
      defer response.Body.Close()

      rawBody, _ := ioutil.ReadAll( response.Body )

      fmt.Println( "Got body", rawBody )

      serializedResponse := t.Serializer.Serialize(response, false).(transports.Response)
      serializedResponse.Body = string(rawBody)

      jsonResponse, _ := json.Marshal(serializedResponse)

      fmt.Println( "Output: ", jsonResponse )

      t.SendMessage( string(jsonResponse) )

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
