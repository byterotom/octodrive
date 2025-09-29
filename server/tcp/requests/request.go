package requests

import (
	"errors"
	"net"
)

const MAX_FILENAME_SIZE = 128

type Request interface {
	HandleConn(conn net.Conn)
}

func NewRequest(conn net.Conn) (Request, error) {
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	var delete, update, send, get byte

	delete = 1 << 0
	update = 1 << 1
	send = 1 << 2
	get = 1 << 3

	req := buf[0] & 15

	switch req {
	case get:
		return &Get{}, nil
	case send:
		return &Send{}, nil
	case update:
		return &Update{}, nil
	case delete:
		return &Delete{}, nil
	default:
	}

	return nil, errors.New("Invalid Request")
}
