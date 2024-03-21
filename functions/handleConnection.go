package functions

import (
	"bufio"
	"errors"
	"net"
)

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
	msg, _, err := bufio.NewReader(conn).ReadLine()
	if err != nil {
		return "", errors.New("403")
	}
	if IsKeys(msg) {
		return "", errors.New(NonValidInputMsg)
	}
	return string(msg), nil
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

func IsKeys(bytes []byte) bool {
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case 0:
			return true
		case 27:
			return true
		}
	}

	return false
}
