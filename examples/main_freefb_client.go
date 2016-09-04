package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/1Conan/transports"
	"os"
)

func main() {
	godotenv.Load()

	fmt.Println("Transports test")

	freefbTransport := transports.FreeFBTransport{
		Login:    os.Getenv("FB_LOGIN"),
		Password: os.Getenv("FB_PASSWORD"),
		Friend:   os.Getenv("FB_FRIEND"),
	}

	Proxy := transports.Proxy{
		Transport: freefbTransport,
		Port:      8080,
	}

	Proxy.Listen()
}
