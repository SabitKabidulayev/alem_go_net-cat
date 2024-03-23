package functions

import (
	"bufio"
	"errors"
	"fmt"
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
	if !IsPrintable(name) {
		return errors.New(LatinNameMsg)
	}
	if len(name) > 12 {
		return errors.New(LongNameMsg)
	}
	if IsUsed(room, name) {
		return errors.New(fmt.Sprintf(TakenNameMsg, name))
	}

	return nil
}

func sendErrorMessage(conn net.Conn, err error) {
	conn.Write([]byte(err.Error()))
}

func (room *Server) AddClient(conn net.Conn, name string) {
	room.mutex.Lock()
	defer room.mutex.Unlock()
	room.clients[conn] = name
}

func (room *Server) informNewConnection(conn net.Conn, ch chan<- Message) {
	msgStruct := Message{false, conn, fmt.Sprintf(NewConnectionMsg, room.clients[conn])}
	ch <- msgStruct
}

func (room *Server) uploadHistory(conn net.Conn) {
	room.mutex.Lock()
	defer room.mutex.Unlock()

	if len(room.history) != 0 {
		for _, msg := range room.history {
			conn.Write([]byte(msg + "\n"))
		}
	}
}

func IsPrintable(msg string) bool {
	printableFlag := false
	for _, char := range msg {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			printableFlag = true
			if char < 32 || char > 126 {
				return false
			}
		}
	}
	return printableFlag
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

func IsUsed(room *Server, str string) bool {
	room.mutex.Lock()
	defer room.mutex.Unlock()

	for _, name := range room.clients {
		if name == str {
			return true
		}
	}
	return false
}
