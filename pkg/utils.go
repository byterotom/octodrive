package pkg

import "crypto/rand"

const CHALLENGE_SIZE = 16

func RandomChallenge() []byte {
	buf := make([]byte, CHALLENGE_SIZE)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}
