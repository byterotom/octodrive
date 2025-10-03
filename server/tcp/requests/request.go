package requests

import (
	"errors"
	"net"
)

const MAX_FILENAME_SIZE = 128

// Request interface for Loose Coupling
type Request interface {
	HandleConn(conn net.Conn)
}

// New request initializer
func NewRequest(conn net.Conn) (Request, error) {
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	var delete, update, upload, download byte
	delete = 1 << 0   // 0000 0001
	update = 1 << 1   // 0000 0010
	upload = 1 << 2   // 0000 0100
	download = 1 << 3 // 0000 1000

	req := buf[0] & 15 // Checking the set bit

	switch req {
	case download:
		return &Download{}, nil
	case upload:
		return &Upload{}, nil
	case update:
		return &Update{}, nil
	case delete:
		return &Delete{}, nil
	default:
	}

	return nil, errors.New("Invalid Request")
}
