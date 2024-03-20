package functions

import "net"

func HandleConnection(room *Server, conn net.Conn, ch chan<- Message) {
	conn.Write([]byte(WelcomeMsg))

	for {
		conn.Write([]byte(EnterNameMsg))

		name, err := ReadLine(conn)
		if err != nil {
			sendErrorMessage(conn, err)
			continue
		}

		if err = room.ValidName(name); err != nil {
			sendErrorMessage(conn, err)
			continue
		}

		room.AddClient(conn, name)

		conn.Write([]byte(JoinChatMsg))

		room.informNewConnection(conn, ch)

		room.uploadHistory(conn)

		break
	}

	handleMessages(room, conn, ch)
}

func ReadLine(conn net.Conn) (string, error) {
	return "msg", nil
}

func (room *Server) ValidName(name string) error {
	return nil
}

func sendErrorMessage(conn net.Conn, err error) {

}

func (room *Server) AddClient(conn net.Conn, name string) {

}

func (room *Server) informNewConnection(conn net.Conn, ch chan<- Message) {

}

func (room *Server) uploadHistory(conn net.Conn) {

}
