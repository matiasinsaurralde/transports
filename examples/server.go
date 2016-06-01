package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/matiasinsaurralde/transports"
	"os"
)

func main() {
	godotenv.Load()

	fmt.Println("Transports test")

	facebookTransport := transports.FacebookTransport{
		Login:    os.Getenv("FB_LOGIN"),
		Password: os.Getenv("FB_PASSWORD"),
		Friend:   os.Getenv("FB_FRIEND"),
	}

	facebookTransport.Listen()
}
