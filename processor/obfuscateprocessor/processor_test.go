package obfuscateprocessor

import (
	"fmt"
	"testing"

	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/base256"
	"github.com/cyrildever/feistel/common/utils/hash"
	"github.com/stretchr/testify/assert"
)

func TestFormatPreservingEncryption(t *testing.T) {
	source := "123123123123"
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)

	obfuscated, err := cipher.Encrypt(source)
	assert.NoError(t, err)
	str := obfuscated.String()
	assert.Equal(t, len([]rune(str)), len(source)) // The source and the obfuscated result have the same number of characters

	readable := base256.ToBase256Readable(obfuscated.Bytes())

	fmt.Println(readable)
	ascii := obfuscated.String(true)

	assert.Equal(t, len(ascii), len(source)) // You must use the `true` argument to the String() method to be sure of that equality in Go (see below)

	assert.Equal(t, len(obfuscated.Bytes()), len(source))
}
