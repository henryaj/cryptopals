package encode

import (
	"encoding/hex"
	"encoding/base64"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	for i := range input {
		b := input[i] ^ xor[i]
		dest[i] = b
	}

	return dest
}

func SingleByteXOR(input []byte) [][]byte {
	var result [][]byte

	for _, xorTarget := range chars {
		xorTargetByte := byte(xorTarget)
		dest := make([]byte, len(input))
		for i := range input {
			dest[i] = input[i] ^ xorTargetByte
		}

		result = append(result, dest)
	}

	return result
}

func RepeatingKeyXOR(input, key []byte) []byte {
	result  := make([]byte, len(input))
	var keyPointer int

	for i := range input {
		result[i] = input[i] ^ key[keyPointer]
		if keyPointer < (len(key)-1) {
			keyPointer++
		} else {
			keyPointer = 0
		}
	}

	return result
}

