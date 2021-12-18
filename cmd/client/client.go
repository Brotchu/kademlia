package main

import (
	"kademlia/connection"
	"log"
)

func main() {
	mesg := connection.Message{
		Payload: []byte("hello"),
	}
	res, err := connection.SendTCP("127.0.0.1:8000", mesg)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println(res)
}
