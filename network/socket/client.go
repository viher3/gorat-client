package socket

import (
	"fmt"
	"net"

	"github.com/viher3/gorat-client/config"
)

func ConnectToServer(cnf *config.Config) (net.Conn, error) {
	fmt.Println("Connecting to the server " + cnf.ServerAddress + " ...")

	conn, err := net.Dial("tcp", cnf.ServerAddress)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return nil, err
	}

	fmt.Println("Connection success.")

	return conn, err
}

func SendMessage(conn net.Conn, message string) error {
	fmt.Println("Sending message: " + message)
	_, err := conn.Write([]byte(message + "\n"))

	if err != nil {
		fmt.Println("Error sending message:", err)
	}

	return err
}

func ReceiveMessage(conn net.Conn) (string, error) {
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func CloseConnection(conn net.Conn) error {
	return conn.Close()
}
