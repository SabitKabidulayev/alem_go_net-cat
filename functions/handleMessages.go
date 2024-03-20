package functions

import (
	"net"
)

type Message struct {
	time    bool
	from    net.Conn
	message string
}

func handleMessages(room *Server, conn net.Conn, ch chan<- Message) {

}
