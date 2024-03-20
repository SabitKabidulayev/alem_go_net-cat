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
	WelcomeMsg = `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|S
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
`
	EnterNameMsg  = "[Enter your name]: "
	TCPChatMsg    = "TCPChat server is listening on port"
	JoinChatMsg   = "You have joined the chat\n"
	FullServerMsg = "Sorry, the chat is full (10/10 connections)\n\n*Press enter to quit and Try to connect later*\n"
)
