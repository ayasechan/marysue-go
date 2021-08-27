package marysue

import (
	"crypto/cipher"
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"
)

type Marysue struct {
	key    [32]byte
	nonce  []byte
	Cipher cipher.AEAD
}

func NewMarysue(password string) *Marysue {
	key := sha256.Sum256([]byte(password))
	aead, _ := chacha20poly1305.NewX(key[:])
	return &Marysue{
		key:    key,
		nonce:  make([]byte, chacha20poly1305.NonceSizeX),
		Cipher: aead,
	}
}

func (m *Marysue) Encrypt(s string) string {
	data := m.Cipher.Seal(nil, m.nonce, []byte(s), nil)
	return ByteArrayToString(data)
}

func (m *Marysue) Decrypt(s string) (string, error) {
	data := StringToByteArray(s)
	data, err := m.Cipher.Open(nil, m.nonce, data, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
