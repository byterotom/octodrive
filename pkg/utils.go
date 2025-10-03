package pkg

import (
	"crypto/rand"
	"os"
	"path/filepath"
)

const CHALLENGE_SIZE = 16

func RandomChallenge() []byte {
	buf := make([]byte, CHALLENGE_SIZE)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}

func OpenFile(filename string) (*os.File, error) {
	dir := filepath.Dir(filename)
	if dir != "" {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
}
