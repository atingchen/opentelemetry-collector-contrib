package obfuscateprocessor

import (
	"fmt"
	"testing"

	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/hash"
	"github.com/stretchr/testify/assert"
)

func TestFormatPreservingEncryption(t *testing.T) {
	source := "my-source-data"
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)

	obfuscated, err := cipher.Encrypt(source)
	assert.NoError(t, err)
	str := obfuscated.String()
	fmt.Println(str)
	assert.Equal(t, len(source), len(str))
}
