package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/matiasinsaurralde/transports"
	"os"
)

func main() {
	godotenv.Load()

	fmt.Println("Transports test (Whatsapp/Yowsup)")

	whatsappTransport := transports.WhatsappTransport{
		Login:             os.Getenv("WA_CLIENT_LOGIN"),
		Password:          os.Getenv("WA_CLIENT_PASSWORD"),
		Contact:           os.Getenv("WA_CLIENT_CONTACT"),
		YowsupWrapperPort: "8888",
	}

	go whatsappTransport.Listen(nil)

	Proxy := transports.Proxy{
		Transport: whatsappTransport,
		Port:      8080,
	}

	Proxy.Listen()
}
