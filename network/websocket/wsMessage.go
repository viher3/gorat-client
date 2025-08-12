package websocket

import "encoding/json"

type WsMessage struct {
	ID      string                 `json:"id"`
	Payload map[string]interface{} `json:"payload"`
}

// NewWsMessage creates a new WebSocket message with the given ID and payload.
// The payload is a map that can hold any data structure.
func NewWsMessage(id string, payload map[string]interface{}) *WsMessage {
	if payload == nil {
		payload = make(map[string]interface{})
	}
	return &WsMessage{
		ID:      id,
		Payload: payload,
	}
}

// ToJSON converts the wsMessage to a JSON string representation.
// Returns the JSON string and any marshaling error.
func (w *WsMessage) ToJSON() (string, error) {
	jsonData, err := json.Marshal(w)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
