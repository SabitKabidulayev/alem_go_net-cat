package functions

import (
	"fmt"
	"net"
	"time"
)

func RunServer(port string) error {
	room := &Server{clients: make(map[net.Conn]string), history: make([]string, 0)}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Println(TCPChatMsg, port)

	ch := make(chan Message)
	go room.Broadcast(ch)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		room.mutex.Lock()
		if len(room.clients) > 9 {
			conn.Write([]byte(FullServerMsg))
			conn.Close()
		}
		room.mutex.Unlock()

		go HandleConnection()
	}
}

func (room *Server) Broadcast(ch <-chan Message) {
	for {
		msg := <-ch
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		text := ""

		room.mutex.Lock()
		if msg.time {
			text = "[" + timestamp + "][" + room.clients[msg.from] + "]:"
		}

		text += msg.message
		room.history = append(room.history, text)

		for conn := range room.clients {
			if conn != msg.from {
				conn.Write([]byte("\n" + text + "\n"))

				conn.Write([]byte("[" + timestamp + "]"))
				conn.Write([]byte("[" + room.clients[conn] + "]:"))
			}
		}
		room.mutex.Unlock()
	}
}
