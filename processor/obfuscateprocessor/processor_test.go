package obfuscateprocessor

import (
	"fmt"
	"github.com/cyrildever/feistel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatPreservingEncryption(t *testing.T) {
	source := "123123123"

	// Encrypt
	cipher := feistel.NewCipher("some-32-byte-long-key-to-be-safe", 10)
	obfuscated, err := cipher.Encrypt(source)
	assert.NoError(t, err)
	fmt.Println(string(obfuscated))

}
