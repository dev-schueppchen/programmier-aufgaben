package main

import (
	"encoding/json"
)

// Event contains the name of the event
// and the data of the event as interface.
type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`

	rawData []byte
}

// EventHandler is the function which
// will handle a specific event, passing
// the struct of the event itself.
type EventHandler func(*Event)

// NewEventFromRaw returns a new event
// which is JSON decoded from the passed
// raw data byte array.
func NewEventFromRaw(raw []byte) (*Event, error) {
	event := new(Event)
	event.rawData = raw
	err := json.Unmarshal(raw, event)
	return event, err
}

// Raw returns the raw JSON formatted data
// of the event as byte array.
func (e *Event) Raw() ([]byte, error) {
	if e.rawData != nil && len(e.rawData) > 0 {
		return e.rawData, nil
	}
	return json.Marshal(e)
}
