package functions

import (
	"net"
	"sync"
)

type Server struct {
	clients map[net.Conn]string
	mutex   sync.Mutex
	history []string
}

// messages
const (
	TCPChatMsg    = "TCPChat server is listening on port"
	FullServerMsg = "Sorry, the chat is full (10/10 connections)\n\n*Press enter to quit and Try to connect later*\n"
)
