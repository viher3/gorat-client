package websocket

import (
	"fmt"

	"github.com/viher3/gorat-client/config"
	"github.com/viher3/gorat-client/system"
)

func ConnectToServer(cnf *config.Config) {
	fmt.Println("Connecting to the server " + cnf.ServerAddress + " ...")

	if cnf.ServerConnectionMode == "websocket" {
		conn, err := Connect(cnf)
		if err != nil {
			fmt.Println("Error connecting to the server:", err)
			return
		}
		defer conn.Close()

		fmt.Println("Connection success.")

		// Send system info upon connection
		systemInfo := system.GetFullInfo()

		// Convert systemInfo to map[string]interface{}
		systemInfoInterface := make(map[string]interface{}, len(systemInfo))
		for k, v := range systemInfo {
			systemInfoInterface[k] = v
		}
		err = SendMessage(conn, *NewWsMessage("system_info", systemInfoInterface))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Infinite loop to read messages
		// TODO: create a commandBus to handle incoming messages
		fmt.Println("Waiting for messages from the server...")

		for {
			msg, err := ReceiveMessage(conn)
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
