package auth

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"os"
	"strings"
)

const SECRET_PHRASE_SIZE = 10

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

func NewSecretPhrase() (string, error) {
	data, err := os.ReadFile("backend/data/words.txt")
	if err != nil {
		return "", err
	}

	words := strings.Split(string(data), "\n")
	totalWords := len(words)

	phrase := make([]string, SECRET_PHRASE_SIZE)

	for i := range SECRET_PHRASE_SIZE {

		idxBig, err := rand.Int(rand.Reader, big.NewInt(int64(totalWords)))
		if err != nil {
			return "", err
		}

		idx := int(idxBig.Int64())
		phrase[i] = words[idx]
	}

	secretPhrase := strings.Join(phrase, " ")

	return secretPhrase, nil
}

func SaveSecretPhraseOnSystem(secretPhrase string) error {
	file, err := os.Create("secret_phrase.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(secretPhrase)
	if err != nil {
		return err
	}

	return nil
}

func LoadSecretPhraseFromSystem() (string, error) {
	buf, err := os.ReadFile("secret_phrase.txt")
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
