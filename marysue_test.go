package marysue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	password  = "marysue-go"
	plainText = "hello"
)

func TestEncrypt(t *testing.T) {
	m := NewMarysue(password)
	text := m.Encrypt(plainText)
	t.Log(text)
}

func TestDecrypt(t *testing.T) {
	cipherText := "莎莉阳千璃·鸢白洁落安·萦莳夏妖幽黛莹·蔷·黛白琉叶爱茜璃"

	m := NewMarysue(password)
	r, err := m.Decrypt(cipherText)
	if err != nil {
		t.Fatal(t)
	}
	assert.Equal(t, r, plainText, "")
}
