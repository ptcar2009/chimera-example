package main

import (
	"log"
	"os"

	"gitlab.inspr.com/chimera/client-golang"
)

func main() {
	log.Println("Environment variables are:")
	log.Println(os.Environ())

	reader, err := client.NewReader()
	if err != nil {
		panic(err)
	}

	for {
		channel, msg, err := reader.ReadMessage()
		log.Printf("Channel is: %s", *channel)
		log.Printf("Message is: %s", msg)
		log.Printf("Error is: %s", err)
		err = reader.Commit()
		if err != nil {
			log.Println(err)
		}
	}
}
