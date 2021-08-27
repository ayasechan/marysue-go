package marysue

import (
	"crypto/cipher"
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"
)

type Marysue struct {
	Key    [32]byte
	nonce  []byte
	cipher cipher.AEAD
}

func NewMarysue(password string) *Marysue {
	key := sha256.Sum256([]byte(password))
	aead, _ := chacha20poly1305.NewX(key[:])
	return &Marysue{
		Key:    key,
		nonce:  make([]byte, chacha20poly1305.NonceSizeX),
		cipher: aead,
	}
}

func (m *Marysue) Encrypt(s string) string {
	data := m.cipher.Seal(nil, m.nonce, []byte(s), nil)
	return ByteArrayToString(data)
}

func (m *Marysue) Decrypt(s string) (string, error) {
	data := StringToByteArray(s)
	data, err := m.cipher.Open(nil, m.nonce, data, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
