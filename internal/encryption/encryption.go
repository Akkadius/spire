package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

type Encrypter struct {
	logger *logrus.Logger
}

func NewEncrypter() *Encrypter {
	return &Encrypter{}
}

func (e *Encrypter) Encrypt(text string, keyString string) string {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		e.logger.Error(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		e.logger.Error(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		e.logger.Error(err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)
}

func (e *Encrypter) Decrypt(encryptedString string, keyString string) string {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		e.logger.Error(err)
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		e.logger.Error(err)
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		e.logger.Error(err)
	}

	return fmt.Sprintf("%s", plaintext)
}
