package websocket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/viher3/gorat-client/config"
)

func Connect(cnf *config.Config) (*websocket.Conn, error) {
	url := "ws://" + cnf.ServerAddress + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SendMessage(conn *websocket.Conn, message WsMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, data)
}

func ReceiveMessage(conn *websocket.Conn) (string, error) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
