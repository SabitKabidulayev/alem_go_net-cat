package functions

import (
	"net"
	"time"
)

type Message struct {
	time    bool
	from    net.Conn
	message string
}

func handleMessages(room *Server, conn net.Conn, ch chan<- Message) {
	for {

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte("[" + timestamp + "]"))
		conn.Write([]byte("[" + room.clients[conn] + "]:"))

		msg, err := ReadLine(conn)
		if err != nil {

			if err.Error() == "403" {

				room.mutex.Lock()
				nameDelete := room.clients[conn]
				delete(room.clients, conn)
				room.mutex.Unlock()
				connMessage := Message{
					time:    false,
					message: nameDelete + " has left our chat...",
					from:    conn,
				}
				ch <- connMessage
				break
			}
			sendErrorMessage(conn, err)
			continue
		}

		if msg == "" {
			continue
		}

		msgStruct := Message{true, conn, msg}
		ch <- msgStruct
	}
}
