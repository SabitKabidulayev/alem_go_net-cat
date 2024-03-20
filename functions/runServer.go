package functions

import (
	"fmt"
	"net"
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

}
