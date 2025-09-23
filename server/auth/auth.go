package auth

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"os"
	"strings"
)

const SECRET_PHRASE_SIZE = 12

type Auth struct {
	SecretPhrase string
	PrivateKey   ed25519.PrivateKey
	PublicKey    ed25519.PublicKey
}

func NewAuth(secretPhrase string) *Auth {
	phraseHash := sha256.Sum256([]byte(secretPhrase))
	privateKey := ed25519.NewKeyFromSeed(phraseHash[:])
	return &Auth{
		SecretPhrase: secretPhrase,
		PrivateKey:   privateKey,
		PublicKey:    privateKey.Public().(ed25519.PublicKey),
	}
}

func NewSecretPhrase() string {
	data, err := os.ReadFile("server/data/words.txt")
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(data), "\n")
	totalWords := len(words)

	phrase := make([]string, SECRET_PHRASE_SIZE)

	for i := range SECRET_PHRASE_SIZE {

		idxBig, err := rand.Int(rand.Reader, big.NewInt(int64(totalWords)))
		if err != nil {
			panic(err)
		}

		idx := int(idxBig.Int64())
		phrase[i] = words[idx]
	}

	secretPhrase := strings.Join(phrase, " ")

	return secretPhrase
}

func SaveSecretPhraseOnSystem(secretPhrase string) {
	file, err := os.Create("server/data/secret_phrase.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(secretPhrase)
	if err != nil {
		panic(err)
	}
}

func LoadSecretPhraseFromSystem() (string, error) {
	buf, err := os.ReadFile("server/data/secret_phrase.txt")
	return string(buf), err
}