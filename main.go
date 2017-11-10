package main

import (
	"flag"
	"os"

	"./lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens to the specified ip address")
	flag.Parse()

	if isHost {
		connectionIP := os.Args[2]
		lib.RunHost(connectionIP)
	} else {
		connectionIP := os.Args[1]
		lib.RunGuest(connectionIP)
	}
}
