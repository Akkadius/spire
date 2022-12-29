package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/sirupsen/logrus"
	"io"
	"math/big"
)

type Encrypter struct {
	logger        *logrus.Logger
	serverconfig  *serverconfig.EQEmuServerConfig
	encryptionKey string
}

func (e *Encrypter) GetEncryptionKey() string {
	return fmt.Sprintf("%v", e.encryptionKey)
}

func (e *Encrypter) SetEncryptionKey(encryptionKey string) {
	e.encryptionKey = encryptionKey
}

func NewEncrypter(
	logger *logrus.Logger,
	serverconfig *serverconfig.EQEmuServerConfig,
) *Encrypter {
	e := &Encrypter{
		logger:       logger,
		serverconfig: serverconfig,
	}

	e.initializeEncryption()
	e.encryptionKey = e.loadEncryptionKey()

	if len(e.encryptionKey) == 0 {
		e.logger.Fatal("Encryption key is invalid")
	}

	return e
}

func (e *Encrypter) Encrypt(text string, keyString string) string {
	key := []byte(keyString)
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
	key := []byte(keyString)
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

func (e *Encrypter) loadEncryptionKey() string {
	if e.serverconfig.Exists() {
		c := e.serverconfig.Get()
		if len(c.Spire.EncryptionKey) != 0 {
			e.logger.Debug("[encryption] Using eqemu server config encryption key")
			return c.Spire.EncryptionKey
		}
	} else if env.IsEnvLoaded() && len(env.Get("APP_KEY", "")) != 0 {
		e.logger.Debug("[encryption] Using [.env] encryption key")
		return env.Get("APP_KEY", "")
	}

	return ""
}

func (e *Encrypter) initializeEncryption() {
	if e.serverconfig.Exists() {
		c := e.serverconfig.Get()
		if len(c.Spire.EncryptionKey) == 0 {
			c.Spire.EncryptionKey = e.generateRandomHash()
			e.logger.Infoln("[encryption] Initialized encryption key in EQEmu server config [spire:encryption_key]")
			e.serverconfig.Save(c)
		}
	} else if env.IsEnvLoaded() && len(env.Get("APP_KEY", "")) == 0 {
		e.logger.Fatal("[encryption] Application key is not defined, it must be set in [.env]")
	}
}

func (e *Encrypter) generateRandomHash() string {
	hash, err := GenerateRandomString(32)
	if err != nil {
		e.logger.Error(err)
	}

	return hash
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
