package main

import "kademlia/connection"

func main() {
	connection.ListenTCP(":8000")
}
