package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptAES(content []byte, key []byte) (data []byte, err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	return gcm.Seal(nonce, nonce, content, nil), nil
}

func DecryptAES(cipherText []byte, key []byte) (text []byte, err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	return gcm.Open(nil, nonce, cipherText, nil)
}
