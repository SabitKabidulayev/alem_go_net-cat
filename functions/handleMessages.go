package functions

import (
	"net"
)

type Message struct {
	time    bool
	from    net.Conn
	message string
}
