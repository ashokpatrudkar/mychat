package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var port = "8080"

func sum(a int, b int) int {
	return a + b
}

// RunHost takes an ip as argument and listens on it for connections
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenerError := net.Listen("tcp", ipAndPort)
	checkForError(listenerError)

	fmt.Println("Listening for connections on: ", ipAndPort)

	connection, acceptError := listener.Accept()
	checkForError(acceptError)
	fmt.Println("New connection accepted")
	for {
		handleHost(connection)
	}
}

// RunGuest takes an ip as argument and listens on it for connections
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	connection, connectionError := net.Dial("tcp", ipAndPort)
	checkForError(connectionError)
	for {
		handleGuest(connection)
	}
}

func checkForError(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
	}
}

func receiveMessage(connection net.Conn) {
	reader := bufio.NewReader(connection)
	message, readError := reader.ReadString('\n')
	checkForError(readError)

	fmt.Println("Message received: ", message)
}

func sendMessage(connection net.Conn) {
	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)

	replyMessage, replyMessageError := replyReader.ReadString('\n')
	checkForError(replyMessageError)

	fmt.Fprint(connection, replyMessage)
}

func handleHost(connection net.Conn) {
	receiveMessage(connection)
	sendMessage(connection)
}

func handleGuest(connection net.Conn) {
	sendMessage(connection)
	receiveMessage(connection)
}
