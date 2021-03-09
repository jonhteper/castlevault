package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"io"
	"os"
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

func Encrypt(m, label string, key *rsa.PublicKey) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key, []byte(m), []byte(label))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(cipherText, label string, key *rsa.PrivateKey) ([]byte, error) {
	ct, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return []byte{}, err
	}

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, key, ct, []byte(label))
}

func OpenPrivateKey(path string) (key *rsa.PrivateKey, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	return jwt.ParseRSAPrivateKeyFromPEM(file)
}

func OpenPublicKey(path string) (key *rsa.PublicKey, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	return jwt.ParseRSAPublicKeyFromPEM(file)
}
