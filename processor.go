package main

import (
	"fmt"
	"strings"

	"github.com/sensimevanidus/repl"
)

const (
	cmdConnect = "connect"
)

type webSocketProcessor struct {
	isConnected bool
	url         *string
}

func newWebSocketProcessor() repl.Processor {
	return &webSocketProcessor{
		isConnected: false,
		url:         nil,
	}
}

// Process establishes a connection to the given URL if the connect command is given. If there is a connection, it
// sends the input to via the websocket.
func (processor *webSocketProcessor) Process(input []byte) (string, error) {
	commands := strings.Fields(string(input))
	if 2 == len(commands) {
		if cmdConnect == commands[0] {
			returnedString := fmt.Sprintf("\u231b\ufe0f connecting to %v", commands[1])
			if err := processor.connect(commands[1]); err != nil {
				returnedString += fmt.Sprintf("\n[error] could not connect to the url %v\n[error] details: %v", commands[1], err.Error())
			} else {
				returnedString += fmt.Sprintf("\n[ok] connection established")
			}
			return returnedString, nil
		} else {
			return fmt.Sprintf("[error] command %v not recognized", commands[0]), nil
		}
	} else if 1 == len(commands) {
		if !processor.isConnected {
			return fmt.Sprintf("[error] websocket connection is not established, please establish the connection first"), nil
		}
	}

	// TODO: Send the input via the websocket
	return fmt.Sprintf("> %v", string(input)), nil
}

func (processor *webSocketProcessor) connect(url string) error {
	// TODO: Establish a real websocket connection
	processor.url = &url
	processor.isConnected = true
	return nil
}