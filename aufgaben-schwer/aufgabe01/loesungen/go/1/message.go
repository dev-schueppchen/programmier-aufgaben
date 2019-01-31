package main

import (
	"encoding/json"
	"errors"
)

// Message contains the name of the sender,
// the actual content and a UNIX millisec
// timestamp.
type Message struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

type messageEvent struct {
	Data *Message `json:"data"`
}

// NewMessageFromEvent creates a new message object
// from an event object.
func NewMessageFromEvent(e *Event) (*Message, error) {
	raw, err := e.Raw()
	if err != nil {
		return nil, err
	}

	// Parsing JSON data to messageEvent object
	msgEvent := new(messageEvent)
	err = json.Unmarshal(raw, msgEvent)
	if err != nil {
		return nil, err
	}
	if msgEvent.Data == nil {
		return nil, errors.New("empty message")
	}

	// returning the actual message object
	return msgEvent.Data, nil
}

// Raw returns the raw JSON formatted data
// of an message object as byte array.
func (m *Message) Raw() ([]byte, error) {
	return json.Marshal(m)
}
