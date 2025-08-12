package network

import (
	"fmt"

	"github.com/viher3/gorat-client/config"
	"github.com/viher3/gorat-client/network/websocket"
)

func ConnectToServer(cnf *config.Config) {
	fmt.Println("Connecting to the server " + cnf.ServerAddress + " ...")

	if cnf.ServerConnectionMode == "websocket" {
		conn, err := websocket.Connect(cnf)
		if err != nil {
			fmt.Println("Error connecting to the server:", err)
			return
		}
		defer conn.Close()

		fmt.Println("Connection success.")

		// TODO: send initial message with system info for client registration in the server
		err = websocket.SendMessage(conn, *websocket.NewWsMessage("client_registration", map[string]interface{}{
			"client_id":   "mylaptop",
			"client_name": "mylaptop Client",
		}))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Infinite loop to read messages
		// TODO: create a commandBus to handle incoming messages
		fmt.Println("Waiting for messages from the server...")

		for {
			msg, err := websocket.ReceiveMessage(conn)
			if err != nil {
				fmt.Println("Error receiving message:", err)
				break // Exit the loop if there's an error (e.g., connection closed)
			}
			fmt.Println("Received message from server:", msg)
		}

	} else {
		fmt.Println("Unsupported connection mode:", cnf.ServerConnectionMode)
	}

}
