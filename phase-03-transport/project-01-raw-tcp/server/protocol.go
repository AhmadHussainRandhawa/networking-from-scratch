package main

import (
	"fmt"
	"strings"
)

type CommandType int

const (
	CommandUnknown CommandType = iota
	CommandHello
	CommandMessage
	CommandQuit
)

type Command struct {
	Type    CommandType
	Payload string
}

func ParseCommand(line string) Command {
	switch {
	case line == "HELLO":
		return Command{
			Type: CommandHello,
		}

	case line == "QUIT":
		return Command{
			Type: CommandQuit,
		}

	case strings.HasPrefix(line, "MESSAGE "):
		return Command{
			Type:    CommandMessage,
			Payload: strings.TrimPrefix(line, "MESSAGE "),
		}

	default:
		return Command{
			Type: CommandUnknown,
		}
	}
}

func ResponseFor(command Command) (string, error) {
	switch command.Type {
	case CommandHello:
		return "WELCOME\n", nil

	case CommandMessage:
		return "ACK\n", nil

	case CommandQuit:
		return "BYE\n", nil

	default:
		return "", fmt.Errorf("unknown command")
	}
}
