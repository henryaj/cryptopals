package encode

import (
	"encoding/hex"
	"encoding/base64"
)

func HexStringToBytes(input string) []byte {
	b, err := hex.DecodeString(input)
	if err != nil {
		panic("unable to decode hex")
	}

	return b
}

func HexBytesToBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func HexBytesToString(hexBytes []byte) string {
	return hex.EncodeToString(hexBytes)
}

func FixedXOR(input, xor []byte) []byte {
	dest := make([]byte, len(input))

	for i, _ := range input {
		b := input[i] ^ xor[i]
		dest[i] = b
	}

	return dest
}

